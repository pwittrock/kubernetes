## *ReplicationControllerStatus v1*

> Example yaml coming soon...



ReplicationControllerStatus represents the current status of a replication controller.

<aside class="notice">
Appears In  <a href="#replicationcontroller-v1">ReplicationController</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
availableReplicas | integer | The number of available replicas (ready for at least minReadySeconds) for this replication controller.
conditions | [ReplicationControllerCondition](#replicationcontrollercondition-v1) array | Represents the latest available observations of a replication controller's current state.
fullyLabeledReplicas | integer | The number of pods that have labels matching the labels of the pod template of the replication controller.
observedGeneration | integer | ObservedGeneration reflects the generation of the most recently observed replication controller.
readyReplicas | integer | The number of ready replicas for this replication controller.
replicas | integer | Replicas is the most recently oberved number of replicas. More info: http://kubernetes.io/docs/user-guide/replication-controller#what-is-a-replication-controller

