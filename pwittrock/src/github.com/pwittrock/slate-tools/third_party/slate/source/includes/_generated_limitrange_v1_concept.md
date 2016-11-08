

-----------
# LimitRange v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | LimitRange







LimitRange sets resource usage limits for each kind of resource in a Namespace.

<aside class="notice">
Appears In <a href="#limitrangelist-v1">LimitRangeList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [LimitRangeSpec](#limitrangespec-v1) | Spec defines the limits enforced. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status


### LimitRangeSpec v1

<aside class="notice">
Appears In <a href="#limitrange-v1">LimitRange</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
limits | [LimitRangeItem](#limitrangeitem-v1) array | Limits is the list of LimitRangeItem objects that are enforced.

### LimitRangeList v1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [LimitRange](#limitrange-v1) array | Items is a list of LimitRange objects. More info: http://releases.k8s.io/HEAD/docs/design/admission_control_limit_range.md
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds




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



create a LimitRange

### HTTP Request

`POST /api/v1/namespaces/{namespace}/limitranges`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [LimitRange](#limitrange-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [LimitRange](#limitrange-v1) | OK


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



replace the specified LimitRange

### HTTP Request

`PUT /api/v1/namespaces/{namespace}/limitranges/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the LimitRange
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [LimitRange](#limitrange-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [LimitRange](#limitrange-v1) | OK


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



partially update the specified LimitRange

### HTTP Request

`PATCH /api/v1/namespaces/{namespace}/limitranges/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the LimitRange
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [LimitRange](#limitrange-v1) | OK


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



delete a LimitRange

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/limitranges/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the LimitRange
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



read the specified LimitRange

### HTTP Request

`GET /api/v1/namespaces/{namespace}/limitranges/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the LimitRange
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
200 | [LimitRange](#limitrange-v1) | OK


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



list or watch objects of kind LimitRange

### HTTP Request

`GET /api/v1/namespaces/{namespace}/limitranges`

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
200 | [LimitRangeList](#limitrangelist-v1) | OK


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



watch changes to an object of kind LimitRange

### HTTP Request

`GET /api/v1/watch/namespaces/{namespace}/limitranges/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the LimitRange
namespace |  | object name and auth scope, such as for teams and projects
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



delete collection of LimitRange

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/limitranges`

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



list or watch objects of kind LimitRange

### HTTP Request

`GET /api/v1/limitranges`

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
200 | [LimitRangeList](#limitrangelist-v1) | OK


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



watch individual changes to a list of LimitRange

### HTTP Request

`GET /api/v1/watch/namespaces/{namespace}/limitranges`

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


## Watch List All Namespaces

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



watch individual changes to a list of LimitRange

### HTTP Request

`GET /api/v1/watch/limitranges`

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




