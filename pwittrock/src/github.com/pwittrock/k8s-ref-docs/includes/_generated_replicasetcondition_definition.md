## *ReplicaSetCondition v1beta1*

> Example yaml coming soon...



ReplicaSetCondition describes the state of a replica set at a certain point.

<aside class="notice">
Appears In  <a href="#replicasetstatus-v1beta1">ReplicaSetStatus</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
status | string | Status of the condition, one of True, False, Unknown.
type | string | Type of replica set condition.
lastProbeTime | [Time](#time-unversioned) | Last time we probed the condition.
lastTransitionTime | [Time](#time-unversioned) | The last time the condition transitioned from one status to another.
message | string | A human readable message indicating details about the transition.
reason | string | The reason for the condition's last transition.

