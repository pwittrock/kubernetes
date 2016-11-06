

-----------
# HorizontalPodAutoscalerList v1beta1

Group        | Version     | Kind
------------ | ---------- | -----------
Extensions | v1beta1 | HorizontalPodAutoscalerList





<aside class="notice">Other api versions of this object exist: <a href="#horizontalpodautoscalerlist-v1">v1</a> </aside>

> Example yaml coming soon...


list of horizontal pod autoscaler objects.



Field        | Schema     | Description
------------ | ---------- | -----------
items | [HorizontalPodAutoscaler](#horizontalpodautoscaler-v1beta1) array | list of horizontal pod autoscaler objects.
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata.





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



watch individual changes to a list of HorizontalPodAutoscaler

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/horizontalpodautoscalers`

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




