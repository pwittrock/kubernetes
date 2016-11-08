

-----------
# PetSetStatus v1alpha1



Group        | Version     | Kind
------------ | ---------- | -----------
Apps | v1alpha1 | PetSetStatus







PetSetStatus represents the current state of a PetSet.

<aside class="notice">
Appears In <a href="#petset-v1alpha1">PetSet</a> </aside>

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



replace status of the specified PetSet

### HTTP Request

`PUT /apis/apps/v1alpha1/namespaces/{namespace}/petsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PetSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [PetSet](#petset-v1alpha1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PetSet](#petset-v1alpha1) | OK


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



partially update status of the specified PetSet

### HTTP Request

`PATCH /apis/apps/v1alpha1/namespaces/{namespace}/petsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PetSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PetSet](#petset-v1alpha1) | OK



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



read status of the specified PetSet

### HTTP Request

`GET /apis/apps/v1alpha1/namespaces/{namespace}/petsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PetSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PetSet](#petset-v1alpha1) | OK




