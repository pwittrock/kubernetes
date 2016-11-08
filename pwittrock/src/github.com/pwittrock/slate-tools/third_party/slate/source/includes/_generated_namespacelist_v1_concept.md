

-----------
# NamespaceList v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | NamespaceList







NamespaceList is a list of Namespaces.



Field        | Schema     | Description
------------ | ---------- | -----------
items | [Namespace](#namespace-v1) array | Items is the list of Namespace objects in the list. More info: http://kubernetes.io/docs/user-guide/namespaces
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



watch individual changes to a list of Namespace

### HTTP Request

`GET /api/v1/watch/namespaces`

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




