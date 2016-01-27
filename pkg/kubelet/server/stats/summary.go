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

	"github.com/golang/glog"
	cadvisorapiv2 "github.com/google/cadvisor/info/v2"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/resource"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/kubelet/cm"
	"k8s.io/kubernetes/pkg/kubelet/dockertools"
	"k8s.io/kubernetes/pkg/kubelet/leaky"
)

// buildSummary aggregates the data passed as arguments and parses it into a *Summary structure
func buildSummary(
	node *api.Node,
	nodeConfig cm.NodeConfig,
	infos map[string]cadvisorapiv2.ContainerInfo,
	rootFsInfo cadvisorapiv2.FsInfo,
	imageFsInfo cadvisorapiv2.FsInfo) (*Summary, error) {

	rootInfo, found := infos["/"]
	if !found || len(rootInfo.Stats) < 1 || latestContainerStats(&rootInfo) == nil {
		return nil, fmt.Errorf("Missing stats for root container")
	}

	rootStats := containerInfoV2ToStats("", &rootInfo, rootFsInfo, imageFsInfo)
	nodeStats := NodeStats{
		NodeName: node.Name,
		CPU:      rootStats.CPU,
		Memory:   rootStats.Memory,
		Network:  containerInfoV2ToNetworkStats(&rootInfo),
		Fs:       getNodeFsStats(rootFsInfo),
	}

	systemContainers := map[string]string{
		SystemContainerKubelet: nodeConfig.KubeletContainerName,
		SystemContainerRuntime: nodeConfig.DockerDaemonContainerName, // TODO: add support for other runtimes
		SystemContainerMisc:    nodeConfig.SystemContainerName,
	}
	for sys, name := range systemContainers {
		if info, ok := infos[name]; ok {
			nodeStats.SystemContainers = append(nodeStats.SystemContainers, containerInfoV2ToStats(sys, &info, rootFsInfo, imageFsInfo))
		}
	}

	summary := Summary{
		Time: unversioned.NewTime(latestContainerStats(&rootInfo).Timestamp),
		Node: nodeStats,
		Pods: buildSummaryPods(infos, rootFsInfo, imageFsInfo),
	}
	return &summary, nil
}

func getNodeFsStats(rootFsInfo cadvisorapiv2.FsInfo) *FsStats {
	return &FsStats{
		AvailableBytes: resource.NewQuantity(int64(rootFsInfo.Available), resource.BinarySI),
		CapacityBytes:  resource.NewQuantity(int64(rootFsInfo.Capacity), resource.BinarySI),
		UsedBytes:      resource.NewQuantity(int64(rootFsInfo.Usage), resource.BinarySI),
	}
}

func latestContainerStats(info *cadvisorapiv2.ContainerInfo) *cadvisorapiv2.ContainerStats {
	stats := info.Stats
	return stats[len(stats)-1]
}

// buildSummaryPods aggregates and returns the container stats in cinfos by the Pod managing the container.
// Containers not managed by a Pod are omitted.
func buildSummaryPods(
	cinfos map[string]cadvisorapiv2.ContainerInfo,
	rootFsInfo cadvisorapiv2.FsInfo,
	imageFsInfo cadvisorapiv2.FsInfo) []PodStats {

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
			stats.Containers = append(stats.Containers, containerInfoV2ToStats(containerName, &cinfo, rootFsInfo, imageFsInfo))
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
	// TODO: Move these out of dockertools and into a general container pkg + make sure rkt is supported
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

func containerInfoV2ToStats(name string, info *cadvisorapiv2.ContainerInfo, rootFsInfo cadvisorapiv2.FsInfo,
	imageFsInfo cadvisorapiv2.FsInfo) ContainerStats {
	stats := ContainerStats{
		Name: name,
	}
	if len(info.Stats) < 1 {
		return stats
	}
	cstat := latestContainerStats(info)
	// TODO(timstclair) - check for overflow
	if info.Spec.HasCpu {
		cpuStats := CPUStats{}
		if cstat.CpuInst != nil {
			cpuStats.UsageCores = resource.NewScaledQuantity(int64(cstat.CpuInst.Usage.Total), resource.Nano)
		}
		if cstat.Cpu != nil {
			cpuStats.UsageCoreSeconds = resource.NewScaledQuantity(int64(cstat.Cpu.Usage.Total), resource.Nano)
		}
		stats.CPU = &cpuStats
	}
	if info.Spec.HasMemory {
		pageFaults := int64(cstat.Memory.ContainerData.Pgfault)
		majorPageFaults := int64(cstat.Memory.ContainerData.Pgmajfault)
		stats.Memory = &MemoryStats{
			UsageBytes:      resource.NewQuantity(int64(cstat.Memory.Usage), resource.DecimalSI),
			WorkingSetBytes: resource.NewQuantity(int64(cstat.Memory.WorkingSet), resource.DecimalSI),
			PageFaults:      &pageFaults,
			MajorPageFaults: &majorPageFaults,
		}
	}
	containerInfoV2FsStats(info, rootFsInfo, imageFsInfo, &stats)
	return stats
}

// containerInfoV2FsStats populates the container fs stats
func containerInfoV2FsStats(
	info *cadvisorapiv2.ContainerInfo,
	rootFsInfo cadvisorapiv2.FsInfo,
	imageFsInfo cadvisorapiv2.FsInfo,
	cs *ContainerStats) {

	cfs := latestContainerStats(info).Filesystem

	glog.Infof("pwittroc - rootfs %+v imageFs %+v info %+v", rootFsInfo, imageFsInfo, cfs)

	// The container logs live on the node rootfs device
	var logsUsage *resource.Quantity
	if cfs.BaseUsageBytes != nil && cfs.TotalUsageBytes != nil {
		logsUsage = resource.NewQuantity(int64(*cfs.TotalUsageBytes - *cfs.BaseUsageBytes), resource.BinarySI)
	}
	cs.Logs = &FsStats{
		AvailableBytes: resource.NewQuantity(int64(rootFsInfo.Available), resource.BinarySI),
		CapacityBytes:  resource.NewQuantity(int64(rootFsInfo.Capacity), resource.BinarySI),
		UsedBytes:      logsUsage,
	}

	// The container rootFs lives on the imageFs devices (which may not be the node root fs)
	var rootUsage *resource.Quantity
	if cfs.BaseUsageBytes != nil {
		rootUsage = resource.NewQuantity(int64(*cfs.BaseUsageBytes), resource.BinarySI)
	}
	cs.Rootfs = &FsStats{
		AvailableBytes: resource.NewQuantity(int64(imageFsInfo.Available), resource.BinarySI),
		CapacityBytes:  resource.NewQuantity(int64(imageFsInfo.Capacity), resource.BinarySI),
		UsedBytes:      rootUsage,
	}
}

func containerInfoV2ToNetworkStats(info *cadvisorapiv2.ContainerInfo) *NetworkStats {
	if !info.Spec.HasNetwork {
		return nil
	}
	if len(info.Stats) < 1 {
		return nil
	}
	var (
		rxBytes  int64
		rxErrors int64
		txBytes  int64
		txErrors int64
	)
	// TODO(stclair): check for overflow
	for _, inter := range latestContainerStats(info).Network.Interfaces {
		rxBytes += int64(inter.RxBytes)
		rxErrors += int64(inter.RxErrors)
		txBytes += int64(inter.TxBytes)
		txErrors += int64(inter.TxErrors)
	}
	return &NetworkStats{
		RxBytes:  resource.NewQuantity(rxBytes, resource.DecimalSI),
		RxErrors: &rxErrors,
		TxBytes:  resource.NewQuantity(txBytes, resource.DecimalSI),
		TxErrors: &txErrors,
	}
}
