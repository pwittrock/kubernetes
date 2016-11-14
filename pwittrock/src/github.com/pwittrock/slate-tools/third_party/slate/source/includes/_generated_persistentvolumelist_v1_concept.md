

-----------
# PersistentVolumeList v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | PersistentVolumeList







PersistentVolumeList is a list of PersistentVolume items.



Field        | Schema     | Description
------------ | ---------- | -----------
items | [PersistentVolume](#persistentvolume-v1) array | List of persistent volumes. More info: http://kubernetes.io/docs/user-guide/persistent-volumes
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds





## <strong>Read Operations</strong>

See supported operations below...

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




