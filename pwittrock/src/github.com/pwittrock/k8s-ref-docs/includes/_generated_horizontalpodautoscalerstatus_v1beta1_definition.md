## *HorizontalPodAutoscalerStatus v1beta1*

> Example yaml coming soon...



current status of a horizontal pod autoscaler



Field        | Schema     | Description
------------ | ---------- | -----------
currentCPUUtilizationPercentage | integer | current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
currentReplicas | integer | current number of replicas of pods managed by this autoscaler.
desiredReplicas | integer | desired number of replicas of pods managed by this autoscaler.
lastScaleTime | [Time](#time-unversioned) | last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
observedGeneration | integer | most recent generation observed by this autoscaler.

