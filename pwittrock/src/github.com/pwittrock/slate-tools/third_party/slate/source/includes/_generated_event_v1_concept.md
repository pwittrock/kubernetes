

-----------
# Event v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | Event




<aside class="notice">Other api versions of this object exist: <a href="#event-versioned">versioned</a> </aside>


Event is a report of an event somewhere in the cluster.

<aside class="notice">
Appears In <a href="#eventlist-v1">EventList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
count | integer | The number of times this event has occurred.
firstTimestamp | [Time](#time-unversioned) | The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
involvedObject | [ObjectReference](#objectreference-v1) | The object that this event is about.
lastTimestamp | [Time](#time-unversioned) | The time at which the most recent occurrence of this event was recorded.
message | string | A human-readable description of the status of this operation.
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
reason | string | This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
source | [EventSource](#eventsource-v1) | The component reporting this event. Should be a short machine understandable string.
type | string | Type of this event (Normal, Warning), new types could be added in the future


### EventList v1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [Event](#event-v1) array | List of events
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



create a Event

### HTTP Request

`POST /api/v1/namespaces/{namespace}/events`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Event](#event-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-v1) | OK


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



replace the specified Event

### HTTP Request

`PUT /api/v1/namespaces/{namespace}/events/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Event
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Event](#event-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-v1) | OK


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



partially update the specified Event

### HTTP Request

`PATCH /api/v1/namespaces/{namespace}/events/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Event
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-v1) | OK


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



delete a Event

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/events/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Event
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



read the specified Event

### HTTP Request

`GET /api/v1/namespaces/{namespace}/events/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Event
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
200 | [Event](#event-v1) | OK


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



list or watch objects of kind Event

### HTTP Request

`GET /api/v1/namespaces/{namespace}/events`

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
200 | [EventList](#eventlist-v1) | OK


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



watch changes to an object of kind Event

### HTTP Request

`GET /api/v1/watch/namespaces/{namespace}/events/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the Event
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



delete collection of Event

### HTTP Request

`DELETE /api/v1/namespaces/{namespace}/events`

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



list or watch objects of kind Event

### HTTP Request

`GET /api/v1/events`

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
200 | [EventList](#eventlist-v1) | OK


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



watch individual changes to a list of Event

### HTTP Request

`GET /api/v1/watch/namespaces/{namespace}/events`

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



watch individual changes to a list of Event

### HTTP Request

`GET /api/v1/watch/events`

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




