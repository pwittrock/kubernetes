## LimitRange v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | LimitRange

> Example yaml coming soon...



LimitRange sets resource usage limits for each kind of resource in a Namespace.

<aside class="notice">
Appears In  <a href="#limitrangelist-v1">LimitRangeList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [LimitRangeSpec](#limitrangespec-v1) | Spec defines the limits enforced. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status

