

-----------
# ReplicationControllerCondition v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ReplicationControllerCondition







> Example yaml coming soon...


ReplicationControllerCondition describes the state of a replication controller at a certain point.

<aside class="notice">
Appears In <a href="#replicationcontrollerstatus-v1">ReplicationControllerStatus</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
lastTransitionTime | [Time](#time-unversioned) | The last time the condition transitioned from one status to another.
message | string | A human readable message indicating details about the transition.
reason | string | The reason for the condition's last transition.
status | string | Status of the condition, one of True, False, Unknown.
type | string | Type of replication controller condition.






