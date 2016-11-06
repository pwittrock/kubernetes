

-----------
# ClusterRoleBindingList v1alpha1

Group        | Version     | Kind
------------ | ---------- | -----------
RbacAuthorization | v1alpha1 | ClusterRoleBindingList







> Example yaml coming soon...


ClusterRoleBindingList is a collection of ClusterRoleBindings



Field        | Schema     | Description
------------ | ---------- | -----------
items | [ClusterRoleBinding](#clusterrolebinding-v1alpha1) array | Items is a list of ClusterRoleBindings
metadata | [ListMeta](#listmeta-unversioned) | Standard object's metadata.





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



watch individual changes to a list of ClusterRoleBinding

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1alpha1/watch/clusterrolebindings`

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




