

-----------
# PodDisruptionBudget v1alpha1

Group        | Version     | Kind
------------ | ---------- | -----------
Policy | v1alpha1 | PodDisruptionBudget







> Example yaml coming soon...


PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods

<aside class="notice">
Appears In <a href="#poddisruptionbudgetlist-v1alpha1">PodDisruptionBudgetList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | 
spec | [PodDisruptionBudgetSpec](#poddisruptionbudgetspec-v1alpha1) | Specification of the desired behavior of the PodDisruptionBudget.
status | [PodDisruptionBudgetStatus](#poddisruptionbudgetstatus-v1alpha1) | Most recently observed status of the PodDisruptionBudget.


### PodDisruptionBudgetSpec v1alpha1

<aside class="notice">
Appears In <a href="#poddisruptionbudget-v1alpha1">PodDisruptionBudget</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
minAvailable | [IntOrString](#intorstring-intstr) | An eviction is allowed if at least "minAvailable" pods selected by "selector" will still be available after the eviction, i.e. even in the absence of the evicted pod.  So for example you can prevent all voluntary evictions by specifying "100%".
selector | [LabelSelector](#labelselector-unversioned) | Label query over pods whose evictions are managed by the disruption budget.

### PodDisruptionBudgetStatus v1alpha1

<aside class="notice">
Appears In <a href="#poddisruptionbudget-v1alpha1">PodDisruptionBudget</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
currentHealthy | integer | current number of healthy pods
desiredHealthy | integer | minimum desired number of healthy pods
disruptionAllowed | boolean | Whether or not a disruption is currently allowed.
expectedPods | integer | total number of pods counted by this disruption budget

### PodDisruptionBudgetList v1alpha1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [PodDisruptionBudget](#poddisruptionbudget-v1alpha1) array | 
metadata | [ListMeta](#listmeta-unversioned) | 




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



create a PodDisruptionBudget

### HTTP Request

`POST /apis/policy/v1alpha1/namespaces/{namespace}/poddisruptionbudgets`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
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



replace the specified PodDisruptionBudget

### HTTP Request

`PUT /apis/policy/v1alpha1/namespaces/{namespace}/poddisruptionbudgets/{name}`

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



partially update the specified PodDisruptionBudget

### HTTP Request

`PATCH /apis/policy/v1alpha1/namespaces/{namespace}/poddisruptionbudgets/{name}`

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



delete a PodDisruptionBudget

### HTTP Request

`DELETE /apis/policy/v1alpha1/namespaces/{namespace}/poddisruptionbudgets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PodDisruptionBudget
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



read the specified PodDisruptionBudget

### HTTP Request

`GET /apis/policy/v1alpha1/namespaces/{namespace}/poddisruptionbudgets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PodDisruptionBudget
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
200 | [PodDisruptionBudget](#poddisruptionbudget-v1alpha1) | OK


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



list or watch objects of kind PodDisruptionBudget

### HTTP Request

`GET /apis/policy/v1alpha1/namespaces/{namespace}/poddisruptionbudgets`

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
200 | [PodDisruptionBudgetList](#poddisruptionbudgetlist-v1alpha1) | OK


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



watch changes to an object of kind PodDisruptionBudget

### HTTP Request

`GET /apis/policy/v1alpha1/watch/namespaces/{namespace}/poddisruptionbudgets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the PodDisruptionBudget
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-versioned) | OK



## <strong>Status & Collection Operations</strong>

See supported operations below...

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



delete collection of PodDisruptionBudget

### HTTP Request

`DELETE /apis/policy/v1alpha1/namespaces/{namespace}/poddisruptionbudgets`

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



list or watch objects of kind PodDisruptionBudget

### HTTP Request

`GET /apis/policy/v1alpha1/poddisruptionbudgets`

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
200 | [PodDisruptionBudgetList](#poddisruptionbudgetlist-v1alpha1) | OK


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



watch individual changes to a list of PodDisruptionBudget

### HTTP Request

`GET /apis/policy/v1alpha1/watch/namespaces/{namespace}/poddisruptionbudgets`

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



watch individual changes to a list of PodDisruptionBudget

### HTTP Request

`GET /apis/policy/v1alpha1/watch/poddisruptionbudgets`

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




