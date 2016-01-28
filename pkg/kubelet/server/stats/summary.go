/*
Copyright 2016 The Kubernetes Authors All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package stats

import (
	"fmt"
	"time"

	"github.com/golang/glog"
	cadvisorapiv1 "github.com/google/cadvisor/info/v1"
	cadvisorapiv2 "github.com/google/cadvisor/info/v2"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/kubelet/cm"
	"k8s.io/kubernetes/pkg/kubelet/dockertools"
	"k8s.io/kubernetes/pkg/kubelet/leaky"
)

// buildSummary aggregates the data passed as arguments and parses it into a *Summary structure
func buildSummary(
	node *api.Node,
	nodeConfig cm.NodeConfig,
	rootFsInfo cadvisorapiv2.FsInfo,
	imageFsInfo cadvisorapiv2.FsInfo,
	infos map[string]cadvisorapiv2.ContainerInfo) (*Summary, error) {

	rootInfo, found := infos["/"]
	if !found {
		return nil, fmt.Errorf("Missing stats for root container")
	}
	cstat, found := latestContainerStats(&rootInfo)
	if !found {
		return nil, fmt.Errorf("Missing stats for root container")
	}

	rootStats := containerInfoV2ToStats("", rootFsInfo, imageFsInfo, &rootInfo)
	nodeStats := NodeStats{
		NodeName: node.Name,
		CPU:      rootStats.CPU,
		Memory:   rootStats.Memory,
		Network:  containerInfoV2ToNetworkStats(&rootInfo),
	}

	systemContainers := map[string]string{
		SystemContainerKubelet: nodeConfig.KubeletContainerName,
		SystemContainerRuntime: nodeConfig.DockerDaemonContainerName, // TODO: add support for other runtimes
		SystemContainerMisc:    nodeConfig.SystemContainerName,
	}
	for sys, name := range systemContainers {
		if info, ok := infos[name]; ok {
			nodeStats.SystemContainers = append(nodeStats.SystemContainers, containerInfoV2ToStats(sys, rootFsInfo, imageFsInfo, &info))
		}
	}

	summary := Summary{
		Time: unversioned.NewTime(cstat.Timestamp),
		Node: nodeStats,
		Pods: buildSummaryPods(rootFsInfo, imageFsInfo, infos),
	}
	return &summary, nil
}

// containerInfoV2FsStats populates the container fs stats
func containerInfoV2FsStats(
	info *cadvisorapiv2.ContainerInfo,
	rootFsInfo cadvisorapiv2.FsInfo,
	imageFsInfo cadvisorapiv2.FsInfo,
	cs *ContainerStats) {

	// The container logs live on the node rootfs device
	cs.Logs = &FsStats{
		AvailableBytes: &rootFsInfo.Available,
		CapacityBytes:  &rootFsInfo.Capacity,
	}

	// The container rootFs lives on the imageFs devices (which may not be the node root fs)
	cs.Rootfs = &FsStats{
		AvailableBytes: &imageFsInfo.Available,
		CapacityBytes:  &imageFsInfo.Capacity,
	}

	lcs, found := latestContainerStats(info)
	if !found {
		return
	}
	cfs := lcs.Filesystem
	if cfs != nil && cfs.BaseUsageBytes != nil {
		// rootfs used == BaseUsageBytes
		cs.Rootfs.UsedBytes = cfs.BaseUsageBytes
		if cfs.TotalUsageBytes != nil {
			// logs used == Total - Base
			logsUsage := *cfs.TotalUsageBytes - *cfs.BaseUsageBytes
			cs.Logs.UsedBytes = &logsUsage
		}
	}
}

func latestContainerStats(info *cadvisorapiv2.ContainerInfo) (*cadvisorapiv2.ContainerStats, bool) {
	stats := info.Stats
	if len(stats) < 1 {
		return nil, false
	}
	latest := stats[len(stats)-1]
	if latest == nil {
		return nil, false
	}
	return latest, true
}

// buildSummaryPods aggregates and returns the container stats in cinfos by the Pod managing the container.
// Containers not managed by a Pod are omitted.
func buildSummaryPods(
	rootFsInfo cadvisorapiv2.FsInfo,
	imageFsInfo cadvisorapiv2.FsInfo,
	cinfos map[string]cadvisorapiv2.ContainerInfo) []PodStats {
	// Map each container to a pod and update the PodStats with container data
	podToStats := map[NonLocalObjectReference]*PodStats{}
	for _, cinfo := range cinfos {
		// Build the Pod key if this container is managed by a Pod
		if !isPodManagedContainer(&cinfo) {
			continue
		}
		ref := buildPodRef(&cinfo)

		// Lookup the PodStats for the pod using the PodRef.  If none exists, initialize a new entry.
		stats, found := podToStats[ref]
		if !found {
			stats = &PodStats{PodRef: ref}
			podToStats[ref] = stats
		}

		// Update the PodStats entry with the stats from the container by adding it to stats.Containers
		containerName := dockertools.GetContainerName(&cinfo)
		if containerName == leaky.PodInfraContainerName {
			// Special case for infrastructure container which is hidden from the user and has network stats
			stats.Network = containerInfoV2ToNetworkStats(&cinfo)
		} else {
			stats.Containers = append(stats.Containers, containerInfoV2ToStats(containerName, rootFsInfo, imageFsInfo, &cinfo))
			stats.UserDefinedMetrics = append(stats.UserDefinedMetrics, containerInfoV2ToUserDefinedMetrics(&cinfo)...)
		}
	}

	// Add each PodStats to the result
	result := make([]PodStats, 0, len(podToStats))
	for _, stats := range podToStats {
		result = append(result, *stats)
	}
	return result
}

// buildPodRef returns a NonLocalObjectReference that identifies the Pod managing cinfo
func buildPodRef(cinfo *cadvisorapiv2.ContainerInfo) NonLocalObjectReference {
	podName := dockertools.GetPodName(cinfo)
	podNamespace := dockertools.GetPodNamespace(cinfo)
	return NonLocalObjectReference{Name: podName, Namespace: podNamespace}
}

// isPodManagedContainer returns true if the cinfo container is managed by a Pod
func isPodManagedContainer(cinfo *cadvisorapiv2.ContainerInfo) bool {
	podName := dockertools.GetPodName(cinfo)
	podNamespace := dockertools.GetPodNamespace(cinfo)
	managed := podName != "" && podNamespace != ""
	if !managed && podName != podNamespace {
		glog.Warningf(
			"Expect container to have either both podName (%s) and podNamespace (%s) labels, or neither.",
			podName, podNamespace)
	}
	return managed
}

func containerInfoV2ToStats(
	name string,
	rootFsInfo cadvisorapiv2.FsInfo,
	imageFsInfo cadvisorapiv2.FsInfo,
	info *cadvisorapiv2.ContainerInfo) ContainerStats {
	stats := ContainerStats{
		Name: name,
	}
	cstat, found := latestContainerStats(info)
	if !found {
		return stats
	}
	// TODO(timstclair) - check for overflow
	if info.Spec.HasCpu {
		cpuStats := CPUStats{}
		if cstat.CpuInst != nil {
			cpuStats.UsageCores = &cstat.CpuInst.Usage.Total
		}
		if cstat.Cpu != nil {
			cpuStats.UsageCoreNanoSeconds = &cstat.Cpu.Usage.Total
		}
		stats.CPU = &cpuStats
	}
	if info.Spec.HasMemory {
		pageFaults := int64(cstat.Memory.ContainerData.Pgfault)
		majorPageFaults := int64(cstat.Memory.ContainerData.Pgmajfault)
		stats.Memory = &MemoryStats{
			UsageBytes:      &cstat.Memory.Usage,
			WorkingSetBytes: &cstat.Memory.WorkingSet,
			PageFaults:      &pageFaults,
			MajorPageFaults: &majorPageFaults,
		}
	}
	containerInfoV2FsStats(info, rootFsInfo, imageFsInfo, &stats)
	return stats
}

func containerInfoV2ToNetworkStats(info *cadvisorapiv2.ContainerInfo) *NetworkStats {
	if !info.Spec.HasNetwork {
		return nil
	}
	cstat, found := latestContainerStats(info)
	if !found {
		return nil
	}
	var (
		rxBytes  uint64
		rxErrors uint64
		txBytes  uint64
		txErrors uint64
	)
	// TODO(stclair): check for overflow
	for _, inter := range cstat.Network.Interfaces {
		rxBytes += inter.RxBytes
		rxErrors += inter.RxErrors
		txBytes += inter.TxBytes
		txErrors += inter.TxErrors
	}
	return &NetworkStats{
		RxBytes:  &rxBytes,
		RxErrors: &rxErrors,
		TxBytes:  &txBytes,
		TxErrors: &txErrors,
	}
}

func containerInfoV2ToUserDefinedMetrics(info *cadvisorapiv2.ContainerInfo) []UserDefinedMetric {
	type specVal struct {
		ref     UserDefinedMetricDescriptor
		valType cadvisorapiv1.DataType
		time    time.Time
		value   float64
	}
	udmMap := map[string]*specVal{}
	for _, spec := range info.Spec.CustomMetrics {
		udmMap[spec.Name] = &specVal{
			ref: UserDefinedMetricDescriptor{
				Name:  spec.Name,
				Type:  UserDefinedMetricType(spec.Type),
				Units: spec.Units,
			},
			valType: spec.Format,
		}
	}
	for _, stat := range info.Stats {
		for name, values := range stat.CustomMetrics {
			specVal, ok := udmMap[name]
			if !ok {
				glog.Warningf("spec for custom metric %q is missing from cAdvisor output. Spec: %+v, Metrics: %+v", name, info.Spec, stat.CustomMetrics)
				continue
			}
			for _, value := range values {
				// Pick the most recent value
				if value.Timestamp.Before(specVal.time) {
					continue
				}
				specVal.time = value.Timestamp
				specVal.value = value.FloatValue
				if specVal.valType == cadvisorapiv1.IntType {
					specVal.value = float64(value.IntValue)
				}
			}
		}
	}
	var udm []UserDefinedMetric
	for _, specVal := range udmMap {
		udm = append(udm, UserDefinedMetric{
			UserDefinedMetricDescriptor: specVal.ref,
			Value: specVal.value,
		})
	}
	return udm
}
