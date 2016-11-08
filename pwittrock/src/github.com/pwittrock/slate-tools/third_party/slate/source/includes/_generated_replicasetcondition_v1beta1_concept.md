

-----------
# ReplicaSetCondition v1beta1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1beta1 | ReplicaSetCondition







ReplicaSetCondition describes the state of a replica set at a certain point.

<aside class="notice">
Appears In <a href="#replicasetstatus-v1beta1">ReplicaSetStatus</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
lastTransitionTime | [Time](#time-unversioned) | The last time the condition transitioned from one status to another.
message | string | A human readable message indicating details about the transition.
reason | string | The reason for the condition's last transition.
status | string | Status of the condition, one of True, False, Unknown.
type | string | Type of replica set condition.






