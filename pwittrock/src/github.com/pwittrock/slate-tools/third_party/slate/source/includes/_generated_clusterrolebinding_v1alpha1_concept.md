

-----------
# ClusterRoleBinding v1alpha1



Group        | Version     | Kind
------------ | ---------- | -----------
RbacAuthorization | v1alpha1 | ClusterRoleBinding







ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.

<aside class="notice">
Appears In <a href="#clusterrolebindinglist-v1alpha1">ClusterRoleBindingList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata.
roleRef | [RoleRef](#roleref-v1alpha1) | RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
subjects | [Subject](#subject-v1alpha1) array | Subjects holds references to the objects the role applies to.


### ClusterRoleBindingList v1alpha1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [ClusterRoleBinding](#clusterrolebinding-v1alpha1) array | Items is a list of ClusterRoleBindings
metadata | [ListMeta](#listmeta-unversioned) | Standard object's metadata.




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



create a ClusterRoleBinding

### HTTP Request

`POST /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ClusterRoleBinding](#clusterrolebinding-v1alpha1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ClusterRoleBinding](#clusterrolebinding-v1alpha1) | OK


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



replace the specified ClusterRoleBinding

### HTTP Request

`PUT /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ClusterRoleBinding
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ClusterRoleBinding](#clusterrolebinding-v1alpha1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ClusterRoleBinding](#clusterrolebinding-v1alpha1) | OK


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



partially update the specified ClusterRoleBinding

### HTTP Request

`PATCH /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ClusterRoleBinding
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ClusterRoleBinding](#clusterrolebinding-v1alpha1) | OK


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



delete a ClusterRoleBinding

### HTTP Request

`DELETE /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ClusterRoleBinding
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



read the specified ClusterRoleBinding

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ClusterRoleBinding
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ClusterRoleBinding](#clusterrolebinding-v1alpha1) | OK


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



list or watch objects of kind ClusterRoleBinding

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings`

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
200 | [ClusterRoleBindingList](#clusterrolebindinglist-v1alpha1) | OK


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



watch changes to an object of kind ClusterRoleBinding

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1alpha1/watch/clusterrolebindings/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the ClusterRoleBinding
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



delete collection of ClusterRoleBinding

### HTTP Request

`DELETE /apis/rbac.authorization.k8s.io/v1alpha1/clusterrolebindings`

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




