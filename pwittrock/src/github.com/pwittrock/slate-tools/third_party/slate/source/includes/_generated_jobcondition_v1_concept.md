

-----------
# JobCondition v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | JobCondition




<aside class="notice">Other api versions of this object exist: <a href="#jobcondition-v1beta1">v1beta1</a> </aside>


JobCondition describes current state of a job.

<aside class="notice">
Appears In <a href="#jobstatus-v1">JobStatus</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
lastProbeTime | [Time](#time-unversioned) | Last time the condition was checked.
lastTransitionTime | [Time](#time-unversioned) | Last time the condition transit from one status to another.
message | string | Human readable message indicating details about last transition.
reason | string | (brief) reason for the condition's last transition.
status | string | Status of the condition, one of True, False, Unknown.
type | string | Type of job condition, Complete or Failed.






