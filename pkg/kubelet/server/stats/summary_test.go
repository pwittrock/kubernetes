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
	"testing"

	"github.com/google/cadvisor/info/v1"
	"github.com/google/cadvisor/info/v2"
	"github.com/stretchr/testify/assert"

	"k8s.io/kubernetes/pkg/api"
	"k8s.io/kubernetes/pkg/api/resource"
	"k8s.io/kubernetes/pkg/kubelet/cm"
	"k8s.io/kubernetes/pkg/kubelet/leaky"
)

const (
	// Offsets from seed value in generated container stats.
	offsetCPUUsageCores = iota
	offsetCPUUsageCoreSeconds
	offsetMemPageFaults
	offsetMemMajorPageFaults
	offsetMemUsageBytes
	offsetMemWorkingSetBytes
	offsetNetRxBytes
	offsetNetRxErrors
	offsetNetTxBytes
	offsetNetTxErrors
)

func TestBuildSummary(t *testing.T) {
	node := api.Node{}
	node.Name = "FooNode"
	nodeConfig := cm.NodeConfig{
		DockerDaemonContainerName: "/docker-daemon",
		SystemContainerName:       "/system",
		KubeletContainerName:      "/kubelet",
	}
	const (
		namespace0 = "test0"
		namespace2 = "test2"
	)
	const (
		seedRoot           = 0
		seedRuntime        = 100
		seedKubelet        = 200
		seedMisc           = 300
		seedPod0Infra      = 1000
		seedPod0Container0 = 2000
		seedPod0Container1 = 2001
		seedPod1Infra      = 3000
		seedPod1Container  = 4000
		seedPod2Infra      = 5000
		seedPod2Container  = 6000
	)
	const (
		pName0 = "pod0"
		pName1 = "pod1"
		pName2 = "pod0" // ensure pName2 conflicts with pName0, but is in a different namespace
	)
	const (
		cName00 = "c0"
		cName01 = "c1"
		cName10 = "c0" // ensure cName10 conflicts with cName02, but is in a different pod
		cName20 = "c1" // ensure cName20 conflicts with cName01, but is in a different pod + namespace
	)

	prf0 := NonLocalObjectReference{Name: pName0, Namespace: namespace0}
	prf1 := NonLocalObjectReference{Name: pName1, Namespace: namespace0}
	prf2 := NonLocalObjectReference{Name: pName2, Namespace: namespace2}
	infos := map[string]v2.ContainerInfo{
		"/":              summaryTestContainerInfo(seedRoot, "", "", ""),
		"/docker-daemon": summaryTestContainerInfo(seedRuntime, "", "", ""),
		"/kubelet":       summaryTestContainerInfo(seedKubelet, "", "", ""),
		"/system":        summaryTestContainerInfo(seedMisc, "", "", ""),
		// Pod0 - Namespace0
		"/pod0-i":  summaryTestContainerInfo(seedPod0Infra, pName0, namespace0, leaky.PodInfraContainerName),
		"/pod0-c0": summaryTestContainerInfo(seedPod0Container0, pName0, namespace0, cName00),
		"/pod0-c2": summaryTestContainerInfo(seedPod0Container1, pName0, namespace0, cName01),
		// Pod1 - Namespace0
		"/pod1-i":  summaryTestContainerInfo(seedPod1Infra, pName1, namespace0, leaky.PodInfraContainerName),
		"/pod1-c0": summaryTestContainerInfo(seedPod1Container, pName1, namespace0, cName10),
		// Pod2 - Namespace2
		"/pod2-i":  summaryTestContainerInfo(seedPod2Infra, pName2, namespace2, leaky.PodInfraContainerName),
		"/pod2-c0": summaryTestContainerInfo(seedPod2Container, pName2, namespace2, cName20),
	}

	summary, err := buildSummary(&node, nodeConfig, infos)
	assert.NoError(t, err)
	nodeStats := summary.Node
	assert.Equal(t, "FooNode", nodeStats.NodeName)
	checkCPUStats(t, "Node", seedRoot, nodeStats.CPU)
	checkMemoryStats(t, "Node", seedRoot, nodeStats.Memory)
	checkNetworkStats(t, "Node", seedRoot, nodeStats.Network)

	systemSeeds := map[string]int{
		SystemContainerRuntime: seedRuntime,
		SystemContainerKubelet: seedKubelet,
		SystemContainerMisc:    seedMisc,
	}
	for _, sys := range nodeStats.SystemContainers {
		name := sys.Name
		seed, found := systemSeeds[name]
		if !found {
			t.Errorf("Unknown SystemContainer: %q", name)
		}
		checkCPUStats(t, name, seed, sys.CPU)
		checkMemoryStats(t, name, seed, sys.Memory)
	}

	assert.Equal(t, 3, len(summary.Pods))
	indexPods := make(map[NonLocalObjectReference]PodStats, len(summary.Pods))
	for _, pod := range summary.Pods {
		indexPods[pod.PodRef] = pod
	}

	// Validate Pod0 Results
	assert.NotNil(t, indexPods[prf0])
	ps := indexPods[prf0]
	assert.Len(t, ps.Containers, 2)
	indexCon := make(map[string]ContainerStats, len(ps.Containers))
	for _, con := range ps.Containers {
		indexCon[con.Name] = con
	}
	con := indexCon[cName00]
	checkCPUStats(t, "container", seedPod0Container0, con.CPU)
	checkMemoryStats(t, "container", seedPod0Container0, con.Memory)

	con = indexCon[cName01]
	checkCPUStats(t, "container", seedPod0Container1, con.CPU)
	checkMemoryStats(t, "container", seedPod0Container1, con.Memory)

	checkNetworkStats(t, "Pod", seedPod0Infra, ps.Network)

	// Validate Pod1 Results
	assert.NotNil(t, indexPods[prf1])
	ps = indexPods[prf1]
	assert.Len(t, ps.Containers, 1)
	con = ps.Containers[0]
	assert.Equal(t, cName10, con.Name)
	checkCPUStats(t, "container", seedPod1Container, con.CPU)
	checkMemoryStats(t, "container", seedPod1Container, con.Memory)
	checkNetworkStats(t, "Pod", seedPod1Infra, ps.Network)

	// Validate Pod2 Results
	assert.NotNil(t, indexPods[prf2])
	ps = indexPods[prf2]
	assert.Len(t, ps.Containers, 1)
	con = ps.Containers[0]
	assert.Equal(t, cName20, con.Name)
	checkCPUStats(t, "container", seedPod2Container, con.CPU)
	checkMemoryStats(t, "container", seedPod2Container, con.Memory)
	checkNetworkStats(t, "Pod", seedPod2Infra, ps.Network)
}

func summaryTestContainerInfo(seed int, podName string, podNamespace string, containerName string) v2.ContainerInfo {
	stats := v2.ContainerStats{
		Cpu:     &v1.CpuStats{},
		CpuInst: &v2.CpuInstStats{},
		Memory: &v1.MemoryStats{
			Usage:      uint64(seed + offsetMemUsageBytes),
			WorkingSet: uint64(seed + offsetMemWorkingSetBytes),
			ContainerData: v1.MemoryStatsMemoryData{
				Pgfault:    uint64(seed + offsetMemPageFaults),
				Pgmajfault: uint64(seed + offsetMemMajorPageFaults),
			},
		},
		Network: &v2.NetworkStats{
			Interfaces: []v1.InterfaceStats{{
				RxBytes:  uint64(seed + offsetNetRxBytes),
				RxErrors: uint64(seed + offsetNetRxErrors),
				TxBytes:  uint64(seed + offsetNetTxBytes),
				TxErrors: uint64(seed + offsetNetTxErrors),
			}},
		},
	}
	stats.Cpu.Usage.Total = uint64(seed + offsetCPUUsageCoreSeconds)
	stats.CpuInst.Usage.Total = uint64(seed + offsetCPUUsageCores)
	labels := map[string]string{}
	if podName != "" {
		labels = map[string]string{
			"io.kubernetes.pod.name":       podName,
			"io.kubernetes.pod.namespace":  podNamespace,
			"io.kubernetes.container.name": containerName,
		}
	}
	return v2.ContainerInfo{
		Spec: v2.ContainerSpec{
			HasCpu:     true,
			HasMemory:  true,
			HasNetwork: true,
			Labels:     labels,
		},
		Stats: []*v2.ContainerStats{&stats},
	}
}

func checkNetworkStats(t *testing.T, label string, seed int, stats *NetworkStats) {
	assert.EqualValues(t, seed+offsetNetRxBytes, stats.RxBytes.Value(), label+".Net.RxBytes")
	assert.EqualValues(t, seed+offsetNetRxErrors, *stats.RxErrors, label+".Net.RxErrors")
	assert.EqualValues(t, seed+offsetNetTxBytes, stats.TxBytes.Value(), label+".Net.TxBytes")
	assert.EqualValues(t, seed+offsetNetTxErrors, *stats.TxErrors, label+".Net.TxErrors")
}

func checkCPUStats(t *testing.T, label string, seed int, stats *CPUStats) {
	assert.EqualValues(t, seed+offsetCPUUsageCores, stats.UsageCores.ScaledValue(resource.Nano), label+".CPU.UsageCores")
	assert.EqualValues(t, seed+offsetCPUUsageCoreSeconds, stats.UsageCoreSeconds.ScaledValue(resource.Nano), label+".CPU.UsageCoreSeconds")
}

func checkMemoryStats(t *testing.T, label string, seed int, stats *MemoryStats) {
	assert.EqualValues(t, seed+offsetMemUsageBytes, stats.UsageBytes.Value(), label+".Mem.UsageBytes")
	assert.EqualValues(t, seed+offsetMemWorkingSetBytes, stats.WorkingSetBytes.Value(), label+".Mem.WorkingSetBytes")
	assert.EqualValues(t, seed+offsetMemPageFaults, *stats.PageFaults, label+".Mem.PageFaults")
	assert.EqualValues(t, seed+offsetMemMajorPageFaults, *stats.MajorPageFaults, label+".Mem.MajorPageFaults")
}
