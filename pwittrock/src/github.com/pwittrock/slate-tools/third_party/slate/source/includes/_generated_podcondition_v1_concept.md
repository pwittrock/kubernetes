

-----------
# PodCondition v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | PodCondition







PodCondition contains details for the current condition of this pod.

<aside class="notice">
Appears In <a href="#podstatus-v1">PodStatus</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
lastProbeTime | [Time](#time-unversioned) | Last time we probed the condition.
lastTransitionTime | [Time](#time-unversioned) | Last time the condition transitioned from one status to another.
message | string | Human-readable message indicating details about last transition.
reason | string | Unique, one-word, CamelCase reason for the condition's last transition.
status | string | Status is the status of the condition. Can be True, False, Unknown. More info: http://kubernetes.io/docs/user-guide/pod-states#pod-conditions
type | string | Type is the type of the condition. Currently only Ready. More info: http://kubernetes.io/docs/user-guide/pod-states#pod-conditions






