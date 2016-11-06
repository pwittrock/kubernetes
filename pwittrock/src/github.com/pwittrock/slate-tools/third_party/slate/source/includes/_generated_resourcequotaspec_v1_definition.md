## ResourceQuotaSpec v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ResourceQuotaSpec

> Example yaml coming soon...



ResourceQuotaSpec defines the desired hard limits to enforce for Quota.

<aside class="notice">
Appears In  <a href="#resourcequota-v1">ResourceQuota</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
hard | object | Hard is the set of desired hard limits for each named resource. More info: http://releases.k8s.io/HEAD/docs/design/admission_control_resource_quota.md#admissioncontrol-plugin-resourcequota
scopes | string array | A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.

