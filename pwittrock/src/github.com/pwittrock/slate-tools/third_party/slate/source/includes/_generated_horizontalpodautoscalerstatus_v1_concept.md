

-----------
# HorizontalPodAutoscalerStatus v1

Group        | Version     | Kind
------------ | ---------- | -----------
Autoscaling | v1 | HorizontalPodAutoscalerStatus





<aside class="notice">Other api versions of this object exist: <a href="#horizontalpodautoscalerstatus-v1beta1">v1beta1</a> </aside>

> Example yaml coming soon...


current status of a horizontal pod autoscaler

<aside class="notice">
Appears In <a href="#horizontalpodautoscaler-v1">HorizontalPodAutoscaler</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
currentCPUUtilizationPercentage | integer | current average CPU utilization over all pods, represented as a percentage of requested CPU, e.g. 70 means that an average pod is using now 70% of its requested CPU.
currentReplicas | integer | current number of replicas of pods managed by this autoscaler.
desiredReplicas | integer | desired number of replicas of pods managed by this autoscaler.
lastScaleTime | [Time](#time-unversioned) | last time the HorizontalPodAutoscaler scaled the number of pods; used by the autoscaler to control how often the number of pods is changed.
observedGeneration | integer | most recent generation observed by this autoscaler.





## <strong>Write Operations</strong>

See supported operations below...

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




