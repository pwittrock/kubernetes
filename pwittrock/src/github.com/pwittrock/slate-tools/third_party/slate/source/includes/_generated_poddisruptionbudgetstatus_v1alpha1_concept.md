

-----------
# PodDisruptionBudgetStatus v1alpha1

Group        | Version     | Kind
------------ | ---------- | -----------
Policy | v1alpha1 | PodDisruptionBudgetStatus







> Example yaml coming soon...


PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.

<aside class="notice">
Appears In <a href="#poddisruptionbudget-v1alpha1">PodDisruptionBudget</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
currentHealthy | integer | current number of healthy pods
desiredHealthy | integer | minimum desired number of healthy pods
disruptionAllowed | boolean | Whether or not a disruption is currently allowed.
expectedPods | integer | total number of pods counted by this disruption budget





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



replace status of the specified PodDisruptionBudget

### HTTP Request

`PUT /apis/policy/v1alpha1/namespaces/{namespace}/poddisruptionbudgets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PodDisruptionBudget
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [PodDisruptionBudget](#poddisruptionbudget-v1alpha1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PodDisruptionBudget](#poddisruptionbudget-v1alpha1) | OK


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



partially update status of the specified PodDisruptionBudget

### HTTP Request

`PATCH /apis/policy/v1alpha1/namespaces/{namespace}/poddisruptionbudgets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PodDisruptionBudget
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PodDisruptionBudget](#poddisruptionbudget-v1alpha1) | OK



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



read status of the specified PodDisruptionBudget

### HTTP Request

`GET /apis/policy/v1alpha1/namespaces/{namespace}/poddisruptionbudgets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PodDisruptionBudget
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PodDisruptionBudget](#poddisruptionbudget-v1alpha1) | OK




