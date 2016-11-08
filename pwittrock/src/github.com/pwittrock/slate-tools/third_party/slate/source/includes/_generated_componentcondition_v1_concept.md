

-----------
# ComponentCondition v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ComponentCondition







Information about the condition of a component.

<aside class="notice">
Appears In <a href="#componentstatus-v1">ComponentStatus</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
error | string | Condition error code for a component. For example, a health check error code.
message | string | Message about the condition for a component. For example, information about a health check.
status | string | Status of the condition for a component. Valid values for "Healthy": "True", "False", or "Unknown".
type | string | Type of condition for a component. Valid value: "Healthy"






