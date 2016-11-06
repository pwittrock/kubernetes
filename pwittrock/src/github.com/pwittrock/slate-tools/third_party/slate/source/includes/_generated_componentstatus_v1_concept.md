

-----------
# ComponentStatus v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ComponentStatus







> Example yaml coming soon...


ComponentStatus (and ComponentStatusList) holds the cluster validation info.

<aside class="notice">
Appears In <a href="#componentstatuslist-v1">ComponentStatusList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
conditions | [ComponentCondition](#componentcondition-v1) array | List of component conditions observed
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata


### ComponentStatusList v1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [ComponentStatus](#componentstatus-v1) array | List of ComponentStatus objects.
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds




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



read the specified ComponentStatus

### HTTP Request

`GET /api/v1/componentstatuses/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ComponentStatus
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ComponentStatus](#componentstatus-v1) | OK


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



list objects of kind ComponentStatus

### HTTP Request

`GET /api/v1/componentstatuses`

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
200 | [ComponentStatusList](#componentstatuslist-v1) | OK




