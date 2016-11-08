## *ResourceQuota v1*

> Example yaml coming soon...



ResourceQuota sets aggregate quota restrictions enforced per namespace

<aside class="notice">
Appears In  <a href="#resourcequotalist-v1">ResourceQuotaList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [ResourceQuotaSpec](#resourcequotaspec-v1) | Spec defines the desired quota. http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status
status | [ResourceQuotaStatus](#resourcequotastatus-v1) | Status defines the actual enforced quota and its current usage. http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status

