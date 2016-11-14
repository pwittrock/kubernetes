

-----------
# PersistentVolumeStatus v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | PersistentVolumeStatus







PersistentVolumeStatus is the current status of a persistent volume.

<aside class="notice">
Appears In <a href="#persistentvolume-v1">PersistentVolume</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
message | string | A human-readable message indicating details about why the volume is in this state.
phase | string | Phase indicates if a volume is available, bound to a claim, or released by a claim. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#phase
reason | string | Reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI.





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




