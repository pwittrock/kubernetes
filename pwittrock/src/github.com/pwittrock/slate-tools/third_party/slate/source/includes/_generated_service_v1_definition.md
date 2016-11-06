## Service v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | Service

> Example yaml coming soon...



Service is a named abstraction of software service (for example, mysql) consisting of local port (for example 3306) that the proxy listens on, and the selector that determines which pods will answer requests sent through the proxy.

<aside class="notice">
Appears In  <a href="#servicelist-v1">ServiceList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [ServiceSpec](#servicespec-v1) | Spec defines the behavior of a service. http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status
status | [ServiceStatus](#servicestatus-v1) | Most recently observed status of the service. Populated by the system. Read-only. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status

