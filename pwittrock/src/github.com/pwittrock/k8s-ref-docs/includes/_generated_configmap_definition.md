## *ConfigMap v1*

> Example yaml coming soon...



ConfigMap holds configuration data for pods to consume.

<aside class="notice">
Appears In  <a href="#configmaplist-v1">ConfigMapList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
data | object | Data contains the configuration data. Each key must be a valid DNS_SUBDOMAIN with an optional leading dot.
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata

