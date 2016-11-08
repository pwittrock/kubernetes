

-----------
# PersistentVolume v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | PersistentVolume


<aside class="notice">These are assigned to <a href="#pod-v1">Pods</a> using <a href="#persistentvolumeclaim-v1">PersistentVolumeClaims</a>.</aside>




PersistentVolume (PV) is a storage resource provisioned by an administrator. It is analogous to a node. More info: http://kubernetes.io/docs/user-guide/persistent-volumes

<aside class="notice">
Appears In <a href="#persistentvolumelist-v1">PersistentVolumeList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [PersistentVolumeSpec](#persistentvolumespec-v1) | Spec defines a specification of a persistent volume owned by the cluster. Provisioned by an administrator. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#persistent-volumes
status | [PersistentVolumeStatus](#persistentvolumestatus-v1) | Status represents the current information/status for the persistent volume. Populated by the system. Read-only. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#persistent-volumes


### PersistentVolumeSpec v1

<aside class="notice">
Appears In <a href="#persistentvolume-v1">PersistentVolume</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
accessModes | string array | AccessModes contains all ways the volume can be mounted. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes
capacity | object | A description of the persistent volume's resources and capacity. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#capacity
claimRef | [ObjectReference](#objectreference-v1) | ClaimRef is part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim. Expected to be non-nil when bound. claim.VolumeName is the authoritative bind between PV and PVC. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#binding
persistentVolumeReclaimPolicy | string | What happens to a persistent volume when released from its claim. Valid options are Retain (default) and Recycle. Recycling must be supported by the volume plugin underlying this persistent volume. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#recycling-policy

### PersistentVolumeStatus v1

<aside class="notice">
Appears In <a href="#persistentvolume-v1">PersistentVolume</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
message | string | A human-readable message indicating details about why the volume is in this state.
phase | string | Phase indicates if a volume is available, bound to a claim, or released by a claim. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#phase
reason | string | Reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI.

### PersistentVolumeList v1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [PersistentVolume](#persistentvolume-v1) array | List of persistent volumes. More info: http://kubernetes.io/docs/user-guide/persistent-volumes
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



create a PersistentVolume

### HTTP Request

`POST /api/v1/persistentvolumes`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [PersistentVolume](#persistentvolume-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolume](#persistentvolume-v1) | OK


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



replace the specified PersistentVolume

### HTTP Request

`PUT /api/v1/persistentvolumes/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolume
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [PersistentVolume](#persistentvolume-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolume](#persistentvolume-v1) | OK


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



partially update the specified PersistentVolume

### HTTP Request

`PATCH /api/v1/persistentvolumes/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolume
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolume](#persistentvolume-v1) | OK


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



delete a PersistentVolume

### HTTP Request

`DELETE /api/v1/persistentvolumes/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolume
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



read the specified PersistentVolume

### HTTP Request

`GET /api/v1/persistentvolumes/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolume
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
exact |  | Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'
export |  | Should this value be exported.  Export strips fields that a user can not specify.

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolume](#persistentvolume-v1) | OK


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



list or watch objects of kind PersistentVolume

### HTTP Request

`GET /api/v1/persistentvolumes`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
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
200 | [PersistentVolumeList](#persistentvolumelist-v1) | OK


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



watch changes to an object of kind PersistentVolume

### HTTP Request

`GET /api/v1/watch/persistentvolumes/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the PersistentVolume
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



partially update status of the specified PersistentVolume

### HTTP Request

`PATCH /api/v1/persistentvolumes/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolume
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolume](#persistentvolume-v1) | OK


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



read status of the specified PersistentVolume

### HTTP Request

`GET /api/v1/persistentvolumes/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolume
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolume](#persistentvolume-v1) | OK


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



replace status of the specified PersistentVolume

### HTTP Request

`PUT /api/v1/persistentvolumes/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the PersistentVolume
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [PersistentVolume](#persistentvolume-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [PersistentVolume](#persistentvolume-v1) | OK


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



delete collection of PersistentVolume

### HTTP Request

`DELETE /api/v1/persistentvolumes`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
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



watch individual changes to a list of PersistentVolume

### HTTP Request

`GET /api/v1/watch/persistentvolumes`

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




