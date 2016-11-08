

-----------
# PetSetList v1alpha1



Group        | Version     | Kind
------------ | ---------- | -----------
Apps | v1alpha1 | PetSetList







PetSetList is a collection of PetSets.



Field        | Schema     | Description
------------ | ---------- | -----------
items | [PetSet](#petset-v1alpha1) array | 
metadata | [ListMeta](#listmeta-unversioned) | 





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



watch individual changes to a list of PetSet

### HTTP Request

`GET /apis/apps/v1alpha1/watch/namespaces/{namespace}/petsets`

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




