

-----------
# PersistentVolumeClaimStatus v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | PersistentVolumeClaimStatus







> Example yaml coming soon...


PersistentVolumeClaimStatus is the current status of a persistent volume claim.

<aside class="notice">
Appears In <a href="#persistentvolumeclaim-v1">PersistentVolumeClaim</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
accessModes | string array | AccessModes contains the actual access modes the volume backing the PVC has. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes-1
capacity | object | Represents the actual resources of the underlying volume.
phase | string | Phase represents the current phase of PersistentVolumeClaim.





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




