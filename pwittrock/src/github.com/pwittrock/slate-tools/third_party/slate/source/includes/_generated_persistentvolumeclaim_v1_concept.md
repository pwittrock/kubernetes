

-----------
# PersistentVolumeClaim v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | PersistentVolumeClaim


<aside class="notice">A <a href="#persistentvolume-v1">PersistentVolume</a> must be allocated in the cluster to use this.</aside>




PersistentVolumeClaim is a user's request for and claim to a persistent volume

<aside class="notice">
Appears In <a href="#persistentvolumeclaimlist-v1">PersistentVolumeClaimList</a> <a href="#petsetspec-v1alpha1">PetSetSpec</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [PersistentVolumeClaimSpec](#persistentvolumeclaimspec-v1) | Spec defines the desired characteristics of a volume requested by a pod author. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#persistentvolumeclaims
status | [PersistentVolumeClaimStatus](#persistentvolumeclaimstatus-v1) | Status represents the current information/status of a persistent volume claim. Read-only. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#persistentvolumeclaims


### PersistentVolumeClaimSpec v1

<aside class="notice">
Appears In <a href="#persistentvolumeclaim-v1">PersistentVolumeClaim</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
accessModes | string array | AccessModes contains the desired access modes the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes-1
resources | [ResourceRequirements](#resourcerequirements-v1) | Resources represents the minimum resources the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#resources
selector | [LabelSelector](#labelselector-unversioned) | A label query over volumes to consider for binding.
volumeName | string | VolumeName is the binding reference to the PersistentVolume backing this claim.

### PersistentVolumeClaimStatus v1

<aside class="notice">
Appears In <a href="#persistentvolumeclaim-v1">PersistentVolumeClaim</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
accessModes | string array | AccessModes contains the actual access modes the volume backing the PVC has. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes-1
capacity | object | Represents the actual resources of the underlying volume.
phase | string | Phase represents the current phase of PersistentVolumeClaim.

### PersistentVolumeClaimList v1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [PersistentVolumeClaim](#persistentvolumeclaim-v1) array | A list of persistent volume claims. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#persistentvolumeclaims
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



create a PersistentVolumeClaim

### HTTP Request

`POST /api/v1/namespaces/{namespace}/persistentvolumeclaims`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [PersistentVolumeClaim](#persistentvolumeclaim-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolumeClaim](#persistentvolumeclaim-v1) | OK


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



replace the specified PersistentVolumeClaim

### HTTP Request

`PUT /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolumeClaim
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [PersistentVolumeClaim](#persistentvolumeclaim-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolumeClaim](#persistentvolumeclaim-v1) | OK


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



partially update the specified PersistentVolumeClaim

### HTTP Request

`PATCH /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolumeClaim
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolumeClaim](#persistentvolumeclaim-v1) | OK


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



delete a PersistentVolumeClaim

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolumeClaim
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



read the specified PersistentVolumeClaim

### HTTP Request

`GET /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolumeClaim
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
200 | [PersistentVolumeClaim](#persistentvolumeclaim-v1) | OK


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



list or watch objects of kind PersistentVolumeClaim

### HTTP Request

`GET /api/v1/namespaces/{namespace}/persistentvolumeclaims`

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
200 | [PersistentVolumeClaimList](#persistentvolumeclaimlist-v1) | OK


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



watch changes to an object of kind PersistentVolumeClaim

### HTTP Request

`GET /api/v1/watch/namespaces/{namespace}/persistentvolumeclaims/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the PersistentVolumeClaim
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



partially update status of the specified PersistentVolumeClaim

### HTTP Request

`PATCH /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolumeClaim
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolumeClaim](#persistentvolumeclaim-v1) | OK


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



read status of the specified PersistentVolumeClaim

### HTTP Request

`GET /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolumeClaim
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolumeClaim](#persistentvolumeclaim-v1) | OK


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



replace status of the specified PersistentVolumeClaim

### HTTP Request

`PUT /api/v1/namespaces/{namespace}/persistentvolumeclaims/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolumeClaim
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [PersistentVolumeClaim](#persistentvolumeclaim-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolumeClaim](#persistentvolumeclaim-v1) | OK


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



delete collection of PersistentVolumeClaim

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/persistentvolumeclaims`

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



list or watch objects of kind PersistentVolumeClaim

### HTTP Request

`GET /api/v1/persistentvolumeclaims`

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
200 | [PersistentVolumeClaimList](#persistentvolumeclaimlist-v1) | OK


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



watch individual changes to a list of PersistentVolumeClaim

### HTTP Request

`GET /api/v1/watch/persistentvolumeclaims`

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




