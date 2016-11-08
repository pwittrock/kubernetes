

-----------

# ConfigMap v1






> Example yaml coming soon...


ConfigMap holds configuration data for pods to consume.

<aside class="notice">
Appears In <a href="#configmaplist-v1">ConfigMapList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
data | object | Data contains the configuration data. Each key must be a valid DNS_SUBDOMAIN with an optional leading dot.
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata


### ConfigMapList v1



Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ListMeta](#listmeta-unversioned) | More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
items | [ConfigMap](#configmap-v1) array | Items is the list of ConfigMaps.




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Examples using curl coming soon...

create a ConfigMap

### HTTP Request

`POST /api/v1/namespaces/{namespace}/configmaps`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ConfigMap](#configmap-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ConfigMap](#configmap-v1) | OK


## Replace

> Examples using curl coming soon...

replace the specified ConfigMap

### HTTP Request

`PUT /api/v1/namespaces/{namespace}/configmaps/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ConfigMap
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ConfigMap](#configmap-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ConfigMap](#configmap-v1) | OK


## Patch

> Examples using curl coming soon...

partially update the specified ConfigMap

### HTTP Request

`PATCH /api/v1/namespaces/{namespace}/configmaps/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ConfigMap
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ConfigMap](#configmap-v1) | OK


## Delete

> Examples using curl coming soon...

delete a ConfigMap

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/configmaps/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ConfigMap
namespace |  | object name and auth scope, such as for teams and projects
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

read the specified ConfigMap

### HTTP Request

`GET /api/v1/namespaces/{namespace}/configmaps/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ConfigMap
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
exact |  | Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'
export |  | Should this value be exported.  Export strips fields that a user can not specify.

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ConfigMap](#configmap-v1) | OK


## List

> Examples using curl coming soon...

list or watch objects of kind ConfigMap

### HTTP Request

`GET /api/v1/namespaces/{namespace}/configmaps`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
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
200 | [ConfigMapList](#configmaplist-v1) | OK


## Watch

> Examples using curl coming soon...

watch changes to an object of kind ConfigMap

### HTTP Request

`GET /api/v1/watch/namespaces/{namespace}/configmaps/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the ConfigMap
namespace |  | object name and auth scope, such as for teams and projects
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

delete collection of ConfigMap

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/configmaps`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
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


## List All Namespaces

> Examples using curl coming soon...

list or watch objects of kind ConfigMap

### HTTP Request

`GET /api/v1/configmaps`

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
200 | [ConfigMapList](#configmaplist-v1) | OK


## Watch List

> Examples using curl coming soon...

watch individual changes to a list of ConfigMap

### HTTP Request

`GET /api/v1/watch/namespaces/{namespace}/configmaps`

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
200 | [Event](#event-v1) | OK


## Watch List All Namespaces

> Examples using curl coming soon...

watch individual changes to a list of ConfigMap

### HTTP Request

`GET /api/v1/watch/configmaps`

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




