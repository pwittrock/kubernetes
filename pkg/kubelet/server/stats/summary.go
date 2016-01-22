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

	cadvisorapiv2 "github.com/google/cadvisor/info/v2"
	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/resource"
	"k8s.io/kubernetes/pkg/api/unversioned"
	"k8s.io/kubernetes/pkg/kubelet/cm"
	"k8s.io/kubernetes/pkg/kubelet/dockertools"
	"k8s.io/kubernetes/pkg/kubelet/leaky"
)

// Handles stats summary requests to /stats/summary
func buildSummary(
	node *api.Node,
	nodeConfig cm.NodeConfig,
	infos map[string]cadvisorapiv2.ContainerInfo) (*Summary, error) {

	rootInfo, found := infos["/"]
	if !found || len(rootInfo.Stats) < 1 || rootInfo.Stats[0] == nil {
		return nil, fmt.Errorf("Missing stats for root container")
	}

	rootStats := containerInfoV2ToStats("", &rootInfo)
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
			nodeStats.SystemContainers = append(nodeStats.SystemContainers, containerInfoV2ToStats(sys, &info))
		}
	}

	summary := Summary{
		Time: unversioned.NewTime(rootInfo.Stats[0].Timestamp),
		Node: nodeStats,
		Pods: buildSummaryPods(infos),
	}
	return &summary, nil
}

// buildSummaryPods aggregates and returns the container stats in cinfos by the Pod managing the container.
// Containers not managed by a Pod are omitted.
func buildSummaryPods(cinfos map[string]cadvisorapiv2.ContainerInfo) []PodStats {
	// Map each container to a pod and update the PodStats with container data
	podToStats := map[NonLocalObjectReference]*PodStats{}
	for _, cinfo := range cinfos {
		// Build the Pod key if this container is managed by a Pod
		if !isPodManagedContainer(&cinfo) {
			continue
		}
		ref := buildPodKey(&cinfo)

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
			stats.Containers = append(stats.Containers, containerInfoV2ToStats(containerName, &cinfo))
		}
	}

	// Add each PodStats to the result
	result := make([]PodStats, 0, len(podToStats))
	for _, stats := range podToStats {
		result = append(result, *stats)
	}
	return result
}

// buildPodKey returns a NonLocalObjectReference that identifies the Pod managing cinfo
func buildPodKey(cinfo *cadvisorapiv2.ContainerInfo) NonLocalObjectReference {
	podName := dockertools.GetPodName(cinfo)
	podNamespace := dockertools.GetPodNamespace(cinfo)
	return NonLocalObjectReference{Name: podName, Namespace: podNamespace}
}

// isPodManagedContainer returns true if the cinfo container is managed by a Pod
func isPodManagedContainer(cinfo *cadvisorapiv2.ContainerInfo) bool {
	return dockertools.GetPodName(cinfo) != "" && dockertools.GetPodNamespace(cinfo) != ""
}

func containerInfoV2ToStats(name string, info *cadvisorapiv2.ContainerInfo) ContainerStats {
	stats := ContainerStats{
		Name: name,
	}
	if len(info.Stats) < 1 {
		return stats
	}
	cstat := info.Stats[0]
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
	return stats
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
	for _, inter := range info.Stats[0].Network.Interfaces {
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
