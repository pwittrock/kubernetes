

-----------

# DaemonSet v1beta1






> Example yaml coming soon...


DaemonSet represents the configuration of a daemon set.

<aside class="notice">
Appears In <a href="#daemonsetlist-v1beta1">DaemonSetList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [DaemonSetSpec](#daemonsetspec-v1beta1) | Spec defines the desired behavior of this daemon set. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status
status | [DaemonSetStatus](#daemonsetstatus-v1beta1) | Status is the current status of this daemon set. This data may be out of date by some window of time. Populated by the system. Read-only. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status


### DaemonSetSpec v1beta1

<aside class="notice">
Appears In <a href="#daemonset-v1beta1">DaemonSet</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
selector | [LabelSelector](#labelselector-v1) | Selector is a label query over pods that are managed by the daemon set. Must match in order to be controlled. If empty, defaulted to labels on Pod template. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
template | [PodTemplateSpec](#podtemplatespec-v1) | Template is the object that describes the pod that will be created. The DaemonSet will create exactly one copy of this pod on every node that matches the template's node selector (or on every node if no node selector is specified). More info: http://kubernetes.io/docs/user-guide/replication-controller#pod-template

### DaemonSetStatus v1beta1

<aside class="notice">
Appears In <a href="#daemonset-v1beta1">DaemonSet</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
currentNumberScheduled | integer | CurrentNumberScheduled is the number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: http://releases.k8s.io/HEAD/docs/admin/daemons.md
desiredNumberScheduled | integer | DesiredNumberScheduled is the total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: http://releases.k8s.io/HEAD/docs/admin/daemons.md
numberMisscheduled | integer | NumberMisscheduled is the number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: http://releases.k8s.io/HEAD/docs/admin/daemons.md
numberReady | integer | NumberReady is the number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.

### DaemonSetList v1beta1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [DaemonSet](#daemonset-v1beta1) array | Items is a list of daemon sets.
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Examples using curl coming soon...

create a DaemonSet

### HTTP Request

`POST /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [DaemonSet](#daemonset-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSet](#daemonset-v1beta1) | OK


## Replace

> Examples using curl coming soon...

replace the specified DaemonSet

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the DaemonSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [DaemonSet](#daemonset-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSet](#daemonset-v1beta1) | OK


## Patch

> Examples using curl coming soon...

partially update the specified DaemonSet

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the DaemonSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSet](#daemonset-v1beta1) | OK


## Delete

> Examples using curl coming soon...

delete a DaemonSet

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the DaemonSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [DeleteOptions](#deleteoptions-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Status](#status-unversioned) | OK



## <strong>Read Operations</strong>

See supported operations below...

## Read

> Examples using curl coming soon...

read the specified DaemonSet

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the DaemonSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
exact |  | Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'
export |  | Should this value be exported.  Export strips fields that a user can not specify.

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSet](#daemonset-v1beta1) | OK


## List

> Examples using curl coming soon...

list or watch objects of kind DaemonSet

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSetList](#daemonsetlist-v1beta1) | OK


## Watch

> Examples using curl coming soon...

watch changes to an object of kind DaemonSet

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/daemonsets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the DaemonSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-v1) | OK



## <strong>Status & Collection Operations</strong>

See supported operations below...

## Patch Status

> Examples using curl coming soon...

partially update status of the specified DaemonSet

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the DaemonSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSet](#daemonset-v1beta1) | OK


## Read Status

> Examples using curl coming soon...

read status of the specified DaemonSet

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the DaemonSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSet](#daemonset-v1beta1) | OK


## Replace Status

> Examples using curl coming soon...

replace status of the specified DaemonSet

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the DaemonSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [DaemonSet](#daemonset-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSet](#daemonset-v1beta1) | OK


## Delete Collection

> Examples using curl coming soon...

delete collection of DaemonSet

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Status](#status-unversioned) | OK


## List All Namespaces

> Examples using curl coming soon...

list or watch objects of kind DaemonSet

### HTTP Request

`GET /apis/extensions/v1beta1/daemonsets`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSetList](#daemonsetlist-v1beta1) | OK


## Watch List

> Examples using curl coming soon...

watch individual changes to a list of DaemonSet

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/daemonsets`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-v1) | OK


## Watch List All Namespaces

> Examples using curl coming soon...

watch individual changes to a list of DaemonSet

### HTTP Request

`GET /apis/extensions/v1beta1/watch/daemonsets`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-v1) | OK




