

-----------

# ClusterRole v1alpha1






> Example yaml coming soon...


ClusterRole is a cluster level, logical grouping of PolicyRules that can be referenced as a unit by a RoleBinding or ClusterRoleBinding.

<aside class="notice">
Appears In <a href="#clusterrolelist-v1alpha1">ClusterRoleList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
rules | [PolicyRule](#policyrule-v1alpha1) array | Rules holds all the PolicyRules for this ClusterRole
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata.


### ClusterRoleList v1alpha1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [ClusterRole](#clusterrole-v1alpha1) array | Items is a list of ClusterRoles
metadata | [ListMeta](#listmeta-unversioned) | Standard object's metadata.




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Examples using curl coming soon...

create a ClusterRole

### HTTP Request

`POST /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ClusterRole](#clusterrole-v1alpha1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ClusterRole](#clusterrole-v1alpha1) | OK


## Replace

> Examples using curl coming soon...

replace the specified ClusterRole

### HTTP Request

`PUT /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ClusterRole
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ClusterRole](#clusterrole-v1alpha1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ClusterRole](#clusterrole-v1alpha1) | OK


## Patch

> Examples using curl coming soon...

partially update the specified ClusterRole

### HTTP Request

`PATCH /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ClusterRole
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ClusterRole](#clusterrole-v1alpha1) | OK


## Delete

> Examples using curl coming soon...

delete a ClusterRole

### HTTP Request

`DELETE /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ClusterRole
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

> Examples using curl coming soon...

read the specified ClusterRole

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ClusterRole
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ClusterRole](#clusterrole-v1alpha1) | OK


## List

> Examples using curl coming soon...

list or watch objects of kind ClusterRole

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles`

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
200 | [ClusterRoleList](#clusterrolelist-v1alpha1) | OK


## Watch

> Examples using curl coming soon...

watch changes to an object of kind ClusterRole

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1alpha1/watch/clusterroles/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the ClusterRole
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-v1) | OK



## <strong>Status & Collection Operations</strong>

See supported operations below...

## Delete Collection

> Examples using curl coming soon...

delete collection of ClusterRole

### HTTP Request

`DELETE /apis/rbac.authorization.k8s.io/v1alpha1/clusterroles`

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

> Examples using curl coming soon...

watch individual changes to a list of ClusterRole

### HTTP Request

`GET /apis/rbac.authorization.k8s.io/v1alpha1/watch/clusterroles`

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
200 | [Event](#event-v1) | OK




