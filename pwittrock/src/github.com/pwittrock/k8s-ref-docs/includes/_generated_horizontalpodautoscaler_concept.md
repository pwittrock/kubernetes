

-----------

# HorizontalPodAutoscaler v1




<aside class="notice">Older api versions of this object exist: <a href="#horizontalpodautoscaler-v1beta1">v1beta1</a> </aside>

> Example yaml coming soon...


configuration of a horizontal pod autoscaler.

<aside class="notice">
Appears In <a href="#horizontalpodautoscalerlist-v1">HorizontalPodAutoscalerList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [HorizontalPodAutoscalerSpec](#horizontalpodautoscalerspec-v1) | behaviour of autoscaler. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status.
status | [HorizontalPodAutoscalerStatus](#horizontalpodautoscalerstatus-v1) | current information about the autoscaler.


### HorizontalPodAutoscalerSpec v1

<aside class="notice">
Appears In <a href="#horizontalpodautoscaler-v1">HorizontalPodAutoscaler</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
maxReplicas | integer | upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
minReplicas | integer | lower limit for the number of pods that can be set by the autoscaler, default 1.
scaleTargetRef | [CrossVersionObjectReference](#crossversionobjectreference-v1) | reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
targetCPUUtilizationPercentage | integer | target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.

### HorizontalPodAutoscalerStatus v1

<aside class="notice">
Appears In <a href="#horizontalpodautoscaler-v1">HorizontalPodAutoscaler</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
lastScaleTime | [Time](#time-unversioned) | last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
observedGeneration | integer | most recent generation observed by this autoscaler.
currentCPUUtilizationPercentage | integer | current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
currentReplicas | integer | current number of replicas of pods managed by this autoscaler.
desiredReplicas | integer | desired number of replicas of pods managed by this autoscaler.

### HorizontalPodAutoscalerList v1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) array | list of horizontal pod autoscaler objects.
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata.




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Examples using curl coming soon...

create a HorizontalPodAutoscaler

### HTTP Request

`POST /apis/extensions/v1beta1/namespaces/{namespace}/horizontalpodautoscalers`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) | OK


## Replace

> Examples using curl coming soon...

replace the specified HorizontalPodAutoscaler

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the HorizontalPodAutoscaler
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) | OK


## Patch

> Examples using curl coming soon...

partially update the specified HorizontalPodAutoscaler

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the HorizontalPodAutoscaler
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) | OK


## Delete

> Examples using curl coming soon...

delete a HorizontalPodAutoscaler

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the HorizontalPodAutoscaler
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

read the specified HorizontalPodAutoscaler

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the HorizontalPodAutoscaler
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
200 | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) | OK


## List

> Examples using curl coming soon...

list or watch objects of kind HorizontalPodAutoscaler

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/horizontalpodautoscalers`

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
200 | [HorizontalPodAutoscalerList](#horizontalpodautoscalerlist-v1) | OK


## Watch

> Examples using curl coming soon...

watch changes to an object of kind HorizontalPodAutoscaler

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/horizontalpodautoscalers/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the HorizontalPodAutoscaler
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

partially update status of the specified HorizontalPodAutoscaler

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the HorizontalPodAutoscaler
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) | OK


## Read Status

> Examples using curl coming soon...

read status of the specified HorizontalPodAutoscaler

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the HorizontalPodAutoscaler
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) | OK


## Replace Status

> Examples using curl coming soon...

replace status of the specified HorizontalPodAutoscaler

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the HorizontalPodAutoscaler
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) | OK


## Delete Collection

> Examples using curl coming soon...

delete collection of HorizontalPodAutoscaler

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/horizontalpodautoscalers`

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

list or watch objects of kind HorizontalPodAutoscaler

### HTTP Request

`GET /apis/extensions/v1beta1/horizontalpodautoscalers`

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
200 | [HorizontalPodAutoscalerList](#horizontalpodautoscalerlist-v1) | OK


## Watch List

> Examples using curl coming soon...

watch individual changes to a list of HorizontalPodAutoscaler

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/horizontalpodautoscalers`

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

watch individual changes to a list of HorizontalPodAutoscaler

### HTTP Request

`GET /apis/extensions/v1beta1/watch/horizontalpodautoscalers`

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




