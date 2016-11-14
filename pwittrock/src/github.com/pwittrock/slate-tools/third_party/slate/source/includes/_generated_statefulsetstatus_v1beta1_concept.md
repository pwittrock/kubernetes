

-----------
# StatefulSetStatus v1beta1



Group        | Version     | Kind
------------ | ---------- | -----------
Apps | v1beta1 | StatefulSetStatus







StatefulSetStatus represents the current state of a StatefulSet.

<aside class="notice">
Appears In <a href="#statefulset-v1beta1">StatefulSet</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
observedGeneration | integer | most recent generation observed by this autoscaler.
replicas | integer | Replicas is the number of actual replicas.





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



replace status of the specified StatefulSet

### HTTP Request

`PUT /apis/apps/v1beta1/namespaces/{namespace}/statefulsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the StatefulSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [StatefulSet](#statefulset-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [StatefulSet](#statefulset-v1beta1) | OK


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



partially update status of the specified StatefulSet

### HTTP Request

`PATCH /apis/apps/v1beta1/namespaces/{namespace}/statefulsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the StatefulSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [StatefulSet](#statefulset-v1beta1) | OK



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



read status of the specified StatefulSet

### HTTP Request

`GET /apis/apps/v1beta1/namespaces/{namespace}/statefulsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the StatefulSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [StatefulSet](#statefulset-v1beta1) | OK




