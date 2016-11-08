

-----------
# LimitRangeItem v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | LimitRangeItem







LimitRangeItem defines a min/max usage limit for any resource that matches on kind.

<aside class="notice">
Appears In <a href="#limitrangespec-v1">LimitRangeSpec</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
default | object | Default resource requirement limit value by resource name if resource limit is omitted.
defaultRequest | object | DefaultRequest is the default resource requirement request value by resource name if resource request is omitted.
max | object | Max usage constraints on this kind by resource name.
maxLimitRequestRatio | object | MaxLimitRequestRatio if specified, the named resource must have a request and limit that are both non-zero where limit divided by request is less than or equal to the enumerated value; this represents the max burst for the named resource.
min | object | Min usage constraints on this kind by resource name.
type | string | Type of resource that this limit applies to.






