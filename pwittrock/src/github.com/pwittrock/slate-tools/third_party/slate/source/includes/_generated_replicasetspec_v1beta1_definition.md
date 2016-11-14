## ReplicaSetSpec v1beta1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1beta1 | ReplicaSetSpec

> Example yaml coming soon...



ReplicaSetSpec is the specification of a ReplicaSet.

<aside class="notice">
Appears In  <a href="#replicaset-v1beta1">ReplicaSet</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
minReadySeconds | integer | Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
replicas | integer | Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: http://kubernetes.io/docs/user-guide/replication-controller#what-is-a-replication-controller
selector | [LabelSelector](#labelselector-unversioned) | Selector is a label query over pods that should match the replica count. If the selector is empty, it is defaulted to the labels present on the pod template. Label keys and values that must match in order to be controlled by this replica set. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
template | [PodTemplateSpec](#podtemplatespec-v1) | Template is the object that describes the pod that will be created if insufficient replicas are detected. More info: http://kubernetes.io/docs/user-guide/replication-controller#pod-template

