

-----------

# ResourceQuota v1






> Example yaml coming soon...


ResourceQuota sets aggregate quota restrictions enforced per namespace

<aside class="notice">
Appears In <a href="#resourcequotalist-v1">ResourceQuotaList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [ResourceQuotaSpec](#resourcequotaspec-v1) | Spec defines the desired quota. http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status
status | [ResourceQuotaStatus](#resourcequotastatus-v1) | Status defines the actual enforced quota and its current usage. http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status


### ResourceQuotaSpec v1

<aside class="notice">
Appears In <a href="#resourcequota-v1">ResourceQuota</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
hard | object | Hard is the set of desired hard limits for each named resource. More info: http://releases.k8s.io/HEAD/docs/design/admission_control_resource_quota.md#admissioncontrol-plugin-resourcequota
scopes | string array | A collection of filters that must match each object tracked by a quota. If not specified, the quota matches all objects.

### ResourceQuotaStatus v1

<aside class="notice">
Appears In <a href="#resourcequota-v1">ResourceQuota</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
hard | object | Hard is the set of enforced hard limits for each named resource. More info: http://releases.k8s.io/HEAD/docs/design/admission_control_resource_quota.md#admissioncontrol-plugin-resourcequota
used | object | Used is the current observed total usage of the resource in the namespace.

### ResourceQuotaList v1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [ResourceQuota](#resourcequota-v1) array | Items is a list of ResourceQuota objects. More info: http://releases.k8s.io/HEAD/docs/design/admission_control_resource_quota.md#admissioncontrol-plugin-resourcequota
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Examples using curl coming soon...

create a ResourceQuota

### HTTP Request

`POST /api/v1/namespaces/{namespace}/resourcequotas`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ResourceQuota](#resourcequota-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ResourceQuota](#resourcequota-v1) | OK


## Replace

> Examples using curl coming soon...

replace the specified ResourceQuota

### HTTP Request

`PUT /api/v1/namespaces/{namespace}/resourcequotas/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ResourceQuota
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ResourceQuota](#resourcequota-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ResourceQuota](#resourcequota-v1) | OK


## Patch

> Examples using curl coming soon...

partially update the specified ResourceQuota

### HTTP Request

`PATCH /api/v1/namespaces/{namespace}/resourcequotas/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ResourceQuota
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ResourceQuota](#resourcequota-v1) | OK


## Delete

> Examples using curl coming soon...

delete a ResourceQuota

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/resourcequotas/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ResourceQuota
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

read the specified ResourceQuota

### HTTP Request

`GET /api/v1/namespaces/{namespace}/resourcequotas/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ResourceQuota
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
200 | [ResourceQuota](#resourcequota-v1) | OK


## List

> Examples using curl coming soon...

list or watch objects of kind ResourceQuota

### HTTP Request

`GET /api/v1/namespaces/{namespace}/resourcequotas`

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
200 | [ResourceQuotaList](#resourcequotalist-v1) | OK


## Watch

> Examples using curl coming soon...

watch changes to an object of kind ResourceQuota

### HTTP Request

`GET /api/v1/watch/namespaces/{namespace}/resourcequotas/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the ResourceQuota
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

## Patch Status

> Examples using curl coming soon...

partially update status of the specified ResourceQuota

### HTTP Request

`PATCH /api/v1/namespaces/{namespace}/resourcequotas/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ResourceQuota
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ResourceQuota](#resourcequota-v1) | OK


## Read Status

> Examples using curl coming soon...

read status of the specified ResourceQuota

### HTTP Request

`GET /api/v1/namespaces/{namespace}/resourcequotas/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ResourceQuota
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ResourceQuota](#resourcequota-v1) | OK


## Replace Status

> Examples using curl coming soon...

replace status of the specified ResourceQuota

### HTTP Request

`PUT /api/v1/namespaces/{namespace}/resourcequotas/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ResourceQuota
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ResourceQuota](#resourcequota-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ResourceQuota](#resourcequota-v1) | OK


## Delete Collection

> Examples using curl coming soon...

delete collection of ResourceQuota

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/resourcequotas`

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

list or watch objects of kind ResourceQuota

### HTTP Request

`GET /api/v1/resourcequotas`

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
200 | [ResourceQuotaList](#resourcequotalist-v1) | OK


## Watch List

> Examples using curl coming soon...

watch individual changes to a list of ResourceQuota

### HTTP Request

`GET /api/v1/watch/namespaces/{namespace}/resourcequotas`

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

watch individual changes to a list of ResourceQuota

### HTTP Request

`GET /api/v1/watch/resourcequotas`

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




