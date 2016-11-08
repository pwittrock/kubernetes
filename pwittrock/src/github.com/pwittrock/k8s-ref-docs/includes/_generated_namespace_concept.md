

-----------

# Namespace v1






> Example yaml coming soon...


Namespace provides a scope for Names. Use of multiple namespaces is optional.

<aside class="notice">
Appears In <a href="#namespacelist-v1">NamespaceList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [NamespaceSpec](#namespacespec-v1) | Spec defines the behavior of the Namespace. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status
status | [NamespaceStatus](#namespacestatus-v1) | Status describes the current status of a Namespace. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status


### NamespaceSpec v1

<aside class="notice">
Appears In <a href="#namespace-v1">Namespace</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
finalizers | string array | Finalizers is an opaque list of values that must be empty to permanently remove object from storage. More info: http://releases.k8s.io/HEAD/docs/design/namespaces.md#finalizers

### NamespaceStatus v1

<aside class="notice">
Appears In <a href="#namespace-v1">Namespace</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
phase | string | Phase is the current lifecycle phase of the namespace. More info: http://releases.k8s.io/HEAD/docs/design/namespaces.md#phases

### NamespaceList v1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [Namespace](#namespace-v1) array | Items is the list of Namespace objects in the list. More info: http://kubernetes.io/docs/user-guide/namespaces
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Examples using curl coming soon...

create a Namespace

### HTTP Request

`POST /api/v1/namespaces`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Namespace](#namespace-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Namespace](#namespace-v1) | OK


## Replace

> Examples using curl coming soon...

replace the specified Namespace

### HTTP Request

`PUT /api/v1/namespaces/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Namespace
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Namespace](#namespace-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Namespace](#namespace-v1) | OK


## Patch

> Examples using curl coming soon...

partially update the specified Namespace

### HTTP Request

`PATCH /api/v1/namespaces/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Namespace
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Namespace](#namespace-v1) | OK


## Delete

> Examples using curl coming soon...

delete a Namespace

### HTTP Request

`DELETE /api/v1/namespaces/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Namespace
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

read the specified Namespace

### HTTP Request

`GET /api/v1/namespaces/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Namespace
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
exact |  | Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'
export |  | Should this value be exported.  Export strips fields that a user can not specify.

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Namespace](#namespace-v1) | OK


## List

> Examples using curl coming soon...

list or watch objects of kind Namespace

### HTTP Request

`GET /api/v1/namespaces`

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
200 | [NamespaceList](#namespacelist-v1) | OK


## Watch

> Examples using curl coming soon...

watch changes to an object of kind Namespace

### HTTP Request

`GET /api/v1/watch/namespaces/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the Namespace
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

## Patch Status

> Examples using curl coming soon...

partially update status of the specified Namespace

### HTTP Request

`PATCH /api/v1/namespaces/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Namespace
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Namespace](#namespace-v1) | OK


## Read Status

> Examples using curl coming soon...

read status of the specified Namespace

### HTTP Request

`GET /api/v1/namespaces/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Namespace
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Namespace](#namespace-v1) | OK


## Replace Status

> Examples using curl coming soon...

replace status of the specified Namespace

### HTTP Request

`PUT /api/v1/namespaces/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Namespace
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Namespace](#namespace-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Namespace](#namespace-v1) | OK


## Delete Collection

> Examples using curl coming soon...

delete collection of Namespace

### HTTP Request

`DELETE /api/v1/namespaces`

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
200 | [Event](#event-v1) | OK




