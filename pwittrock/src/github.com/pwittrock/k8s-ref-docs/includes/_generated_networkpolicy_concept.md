

-----------

# NetworkPolicy v1beta1






> Example yaml coming soon...




<aside class="notice">
Appears In <a href="#networkpolicylist-v1beta1">NetworkPolicyList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [NetworkPolicySpec](#networkpolicyspec-v1beta1) | Specification of the desired behavior for this NetworkPolicy.


### NetworkPolicySpec v1beta1

<aside class="notice">
Appears In <a href="#networkpolicy-v1beta1">NetworkPolicy</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
ingress | [NetworkPolicyIngressRule](#networkpolicyingressrule-v1beta1) array | List of ingress rules to be applied to the selected pods. Traffic is allowed to a pod if namespace.networkPolicy.ingress.isolation is undefined and cluster policy allows it, OR if the traffic source is the pod's local node, OR if the traffic matches at least one ingress rule across all of the NetworkPolicy objects whose podSelector matches the pod. If this field is empty then this NetworkPolicy does not affect ingress isolation. If this field is present and contains at least one rule, this policy allows any traffic which matches at least one of the ingress rules in this list.
podSelector | [LabelSelector](#labelselector-v1) | Selects the pods to which this NetworkPolicy object applies.  The array of ingress rules is applied to any pods selected by this field. Multiple network policies can select the same set of pods.  In this case, the ingress rules for each are combined additively. This field is NOT optional and follows standard label selector semantics. An empty podSelector matches all pods in this namespace.

### NetworkPolicyList v1beta1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [NetworkPolicy](#networkpolicy-v1beta1) array | Items is a list of schema objects.
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Examples using curl coming soon...

create a NetworkPolicy

### HTTP Request

`POST /apis/extensions/v1beta1/namespaces/{namespace}/networkpolicies`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [NetworkPolicy](#networkpolicy-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [NetworkPolicy](#networkpolicy-v1beta1) | OK


## Replace

> Examples using curl coming soon...

replace the specified NetworkPolicy

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/networkpolicies/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the NetworkPolicy
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [NetworkPolicy](#networkpolicy-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [NetworkPolicy](#networkpolicy-v1beta1) | OK


## Patch

> Examples using curl coming soon...

partially update the specified NetworkPolicy

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/networkpolicies/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the NetworkPolicy
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [NetworkPolicy](#networkpolicy-v1beta1) | OK


## Delete

> Examples using curl coming soon...

delete a NetworkPolicy

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/networkpolicies/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the NetworkPolicy
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

read the specified NetworkPolicy

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/networkpolicies/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the NetworkPolicy
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
200 | [NetworkPolicy](#networkpolicy-v1beta1) | OK


## List

> Examples using curl coming soon...

list or watch objects of kind NetworkPolicy

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/networkpolicies`

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
200 | [NetworkPolicyList](#networkpolicylist-v1beta1) | OK


## Watch

> Examples using curl coming soon...

watch changes to an object of kind NetworkPolicy

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/networkpolicies/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the NetworkPolicy
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

delete collection of NetworkPolicy

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/networkpolicies`

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

list or watch objects of kind NetworkPolicy

### HTTP Request

`GET /apis/extensions/v1beta1/networkpolicies`

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
200 | [NetworkPolicyList](#networkpolicylist-v1beta1) | OK


## Watch List

> Examples using curl coming soon...

watch individual changes to a list of NetworkPolicy

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/networkpolicies`

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

watch individual changes to a list of NetworkPolicy

### HTTP Request

`GET /apis/extensions/v1beta1/watch/networkpolicies`

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




