

-----------
# IngressList v1beta1



Group        | Version     | Kind
------------ | ---------- | -----------
Extensions | v1beta1 | IngressList







IngressList is a collection of Ingress.



Field        | Schema     | Description
------------ | ---------- | -----------
items | [Ingress](#ingress-v1beta1) array | Items is the list of Ingress.
metadata | [ListMeta](#listmeta-unversioned) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata





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



watch individual changes to a list of Ingress

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/ingresses`

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




