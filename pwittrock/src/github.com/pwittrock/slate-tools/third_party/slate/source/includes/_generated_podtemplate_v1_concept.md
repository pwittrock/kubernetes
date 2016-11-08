

-----------
# PodTemplate v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | PodTemplate







PodTemplate describes a template for creating copies of a predefined pod.

<aside class="notice">
Appears In <a href="#podtemplatelist-v1">PodTemplateList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
template | [PodTemplateSpec](#podtemplatespec-v1) | Template defines the pods that will be created from this pod template. http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status


### PodTemplateSpec v1

<aside class="notice">
Appears In <a href="#daemonsetspec-v1beta1">DaemonSetSpec</a> <a href="#deploymentspec-v1beta1">DeploymentSpec</a> <a href="#jobspec-v1">JobSpec</a> <a href="#petsetspec-v1alpha1">PetSetSpec</a> <a href="#podtemplate-v1">PodTemplate</a> <a href="#replicasetspec-v1beta1">ReplicaSetSpec</a> <a href="#replicationcontrollerspec-v1">ReplicationControllerSpec</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [PodSpec](#podspec-v1) | Specification of the desired behavior of the pod. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status

### PodTemplateList v1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [PodTemplate](#podtemplate-v1) array | List of pod templates
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds




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



create a PodTemplate

### HTTP Request

`POST /api/v1/namespaces/{namespace}/podtemplates`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [PodTemplate](#podtemplate-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PodTemplate](#podtemplate-v1) | OK


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



replace the specified PodTemplate

### HTTP Request

`PUT /api/v1/namespaces/{namespace}/podtemplates/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PodTemplate
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [PodTemplate](#podtemplate-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PodTemplate](#podtemplate-v1) | OK


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



partially update the specified PodTemplate

### HTTP Request

`PATCH /api/v1/namespaces/{namespace}/podtemplates/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PodTemplate
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PodTemplate](#podtemplate-v1) | OK


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



delete a PodTemplate

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/podtemplates/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PodTemplate
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



read the specified PodTemplate

### HTTP Request

`GET /api/v1/namespaces/{namespace}/podtemplates/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PodTemplate
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
200 | [PodTemplate](#podtemplate-v1) | OK


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



list or watch objects of kind PodTemplate

### HTTP Request

`GET /api/v1/namespaces/{namespace}/podtemplates`

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
200 | [PodTemplateList](#podtemplatelist-v1) | OK


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



watch changes to an object of kind PodTemplate

### HTTP Request

`GET /api/v1/watch/namespaces/{namespace}/podtemplates/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the PodTemplate
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



delete collection of PodTemplate

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/podtemplates`

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



list or watch objects of kind PodTemplate

### HTTP Request

`GET /api/v1/podtemplates`

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
200 | [PodTemplateList](#podtemplatelist-v1) | OK


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



watch individual changes to a list of PodTemplate

### HTTP Request

`GET /api/v1/watch/podtemplates`

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




