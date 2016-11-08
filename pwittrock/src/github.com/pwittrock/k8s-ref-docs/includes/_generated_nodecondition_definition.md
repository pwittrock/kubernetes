## *NodeCondition v1*

> Example yaml coming soon...



NodeCondition contains condition information for a node.

<aside class="notice">
Appears In  <a href="#nodestatus-v1">NodeStatus</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
message | string | Human readable message indicating details about last transition.
reason | string | (brief) reason for the condition's last transition.
status | string | Status of the condition, one of True, False, Unknown.
type | string | Type of node condition.
lastHeartbeatTime | [Time](#time-unversioned) | Last time we got an update on a given condition.
lastTransitionTime | [Time](#time-unversioned) | Last time the condition transit from one status to another.

