## *Pod v1*

> Example yaml coming soon...



Pod is a collection of containers that can run on a host. This resource is created by clients and scheduled onto hosts.

<aside class="notice">
Appears In  <a href="#podlist-v1">PodList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [PodSpec](#podspec-v1) | Specification of the desired behavior of the pod. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status
status | [PodStatus](#podstatus-v1) | Most recently observed status of the pod. This data may not be up to date. Populated by the system. Read-only. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status

