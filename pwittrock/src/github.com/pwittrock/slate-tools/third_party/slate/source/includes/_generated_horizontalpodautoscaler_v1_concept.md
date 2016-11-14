

-----------
# HorizontalPodAutoscaler v1



Group        | Version     | Kind
------------ | ---------- | -----------
Autoscaling | v1 | HorizontalPodAutoscaler




<aside class="notice">Other api versions of this object exist: <a href="#horizontalpodautoscaler-v1beta1">v1beta1</a> </aside>


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
currentCPUUtilizationPercentage | integer | current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
currentReplicas | integer | current number of replicas of pods managed by this autoscaler.
desiredReplicas | integer | desired number of replicas of pods managed by this autoscaler.
lastScaleTime | [Time](#time-unversioned) | last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
observedGeneration | integer | most recent generation observed by this autoscaler.

### HorizontalPodAutoscalerList v1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1) array | list of horizontal pod autoscaler objects.
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata.




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



create a HorizontalPodAutoscaler

### HTTP Request

`POST /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers`

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

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



replace the specified HorizontalPodAutoscaler

### HTTP Request

`PUT /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}`

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

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



partially update the specified HorizontalPodAutoscaler

### HTTP Request

`PATCH /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}`

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

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



delete a HorizontalPodAutoscaler

### HTTP Request

`DELETE /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}`

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
gracePeriodSeconds |  | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.
orphanDependents |  | Should the dependent objects be orphaned. If true/false, the "orphan" finalizer will be added to/removed from the object's finalizers list.

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Status](#status-unversioned) | OK


## Delete Collection

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



delete collection of HorizontalPodAutoscaler

### HTTP Request

`DELETE /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers`

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



## <strong>Read Operations</strong>

See supported operations below...

## Read

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



read the specified HorizontalPodAutoscaler

### HTTP Request

`GET /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}`

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

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



list or watch objects of kind HorizontalPodAutoscaler

### HTTP Request

`GET /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers`

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


## List All Namespaces

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



list or watch objects of kind HorizontalPodAutoscaler

### HTTP Request

`GET /apis/autoscaling/v1/horizontalpodautoscalers`

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


## Watch

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



watch changes to an object of kind HorizontalPodAutoscaler

### HTTP Request

`GET /apis/autoscaling/v1/watch/namespaces/{namespace}/horizontalpodautoscalers/{name}`

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
200 | [Event](#event-versioned) | OK


## Watch List

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



watch individual changes to a list of HorizontalPodAutoscaler

### HTTP Request

`GET /apis/autoscaling/v1/watch/namespaces/{namespace}/horizontalpodautoscalers`

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
200 | [Event](#event-versioned) | OK


## Watch List All Namespaces

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



watch individual changes to a list of HorizontalPodAutoscaler

### HTTP Request

`GET /apis/autoscaling/v1/watch/horizontalpodautoscalers`

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
200 | [Event](#event-versioned) | OK



## <strong>Status Operations</strong>

See supported operations below...

## Patch Status

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



partially update status of the specified HorizontalPodAutoscaler

### HTTP Request

`PATCH /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status`

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

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



read status of the specified HorizontalPodAutoscaler

### HTTP Request

`GET /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status`

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

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



replace status of the specified HorizontalPodAutoscaler

### HTTP Request

`PUT /apis/autoscaling/v1/namespaces/{namespace}/horizontalpodautoscalers/{name}/status`

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




