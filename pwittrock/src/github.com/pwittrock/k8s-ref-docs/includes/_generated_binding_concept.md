

-----------

# Binding v1






> Example yaml coming soon...


Binding ties one object to another. For example, a pod is bound to a node by a scheduler.



Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
target | [ObjectReference](#objectreference-v1) | The target object that you want to bind to the standard object.





## <strong>Write Operations</strong>

See supported operations below...

## Create

> Examples using curl coming soon...

create a Binding

### HTTP Request

`POST /api/v1/namespaces/{namespace}/bindings`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Binding](#binding-v1) | 
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Binding](#binding-v1) | OK




