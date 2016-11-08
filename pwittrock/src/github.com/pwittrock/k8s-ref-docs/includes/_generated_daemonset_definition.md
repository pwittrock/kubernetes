## *DaemonSet v1beta1*

> Example yaml coming soon...



DaemonSet represents the configuration of a daemon set.

<aside class="notice">
Appears In  <a href="#daemonsetlist-v1beta1">DaemonSetList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [DaemonSetSpec](#daemonsetspec-v1beta1) | Spec defines the desired behavior of this daemon set. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status
status | [DaemonSetStatus](#daemonsetstatus-v1beta1) | Status is the current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status

