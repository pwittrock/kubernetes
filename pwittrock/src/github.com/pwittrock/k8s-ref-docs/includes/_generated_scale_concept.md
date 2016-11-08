

-----------

# Scale v1




<aside class="notice">Older api versions of this object exist: <a href="#scale-v1beta1">v1beta1</a> </aside>

> Example yaml coming soon...


Scale represents a scaling request for a resource.



Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object metadata; More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata.
spec | [ScaleSpec](#scalespec-v1) | defines the behavior of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status.
status | [ScaleStatus](#scalestatus-v1) | current status of the scale. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status. Read-only.


### ScaleSpec v1

<aside class="notice">
Appears In <a href="#scale-v1">Scale</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
replicas | integer | desired number of instances for the scaled object.

### ScaleStatus v1

<aside class="notice">
Appears In <a href="#scale-v1">Scale</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
replicas | integer | actual number of observed instances of the scaled object.
selector | string | label query over pods that should match the replicas count. This is same as the label selector but in the string format to avoid introspection by clients. The string will be in the same format as the query-param syntax. More info about label selectors: http://kubernetes.io/docs/user-guide/labels#label-selectors




## <strong>Misc Operations</strong>

See supported operations below...

## Read Scale

> Examples using curl coming soon...

read scale of the specified Scale

### HTTP Request

`GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/scale`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Scale
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Scale](#scale-v1) | OK


## Replace Scale

> Examples using curl coming soon...

replace scale of the specified Scale

### HTTP Request

`PUT /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/scale`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Scale
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Scale](#scale-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Scale](#scale-v1) | OK


## Patch Scale

> Examples using curl coming soon...

partially update scale of the specified Scale

### HTTP Request

`PATCH /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/scale`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Scale
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Scale](#scale-v1) | OK




