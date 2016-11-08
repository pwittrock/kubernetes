

-----------

# Ingress v1beta1






> Example yaml coming soon...


Ingress is a collection of rules that allow inbound connections to reach the endpoints defined by a backend. An Ingress can be configured to give services externally-reachable urls, load balance traffic, terminate SSL, offer name based virtual hosting etc.

<aside class="notice">
Appears In <a href="#ingresslist-v1beta1">IngressList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [IngressSpec](#ingressspec-v1beta1) | Spec is the desired state of the Ingress. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status
status | [IngressStatus](#ingressstatus-v1beta1) | Status is the current state of the Ingress. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status


### IngressSpec v1beta1

<aside class="notice">
Appears In <a href="#ingress-v1beta1">Ingress</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
tls | [IngressTLS](#ingresstls-v1beta1) array | TLS configuration. Currently the Ingress only supports a single TLS port, 443. If multiple members of this list specify different hosts, they will be multiplexed on the same port according to the hostname specified through the SNI TLS extension, if the ingress controller fulfilling the ingress supports SNI.
backend | [IngressBackend](#ingressbackend-v1beta1) | A default backend capable of servicing requests that don't match any rule. At least one of 'backend' or 'rules' must be specified. This field is optional to allow the loadbalancer controller or defaulting logic to specify a global default.
rules | [IngressRule](#ingressrule-v1beta1) array | A list of host rules used to configure the Ingress. If unspecified, or no rule matches, all traffic is sent to the default backend.

### IngressStatus v1beta1

<aside class="notice">
Appears In <a href="#ingress-v1beta1">Ingress</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
loadBalancer | [LoadBalancerStatus](#loadbalancerstatus-v1) | LoadBalancer contains the current status of the load-balancer.

### IngressList v1beta1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [Ingress](#ingress-v1beta1) array | Items is the list of Ingress.
metadata | [ListMeta](#listmeta-unversioned) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Examples using curl coming soon...

create a Ingress

### HTTP Request

`POST /apis/extensions/v1beta1/namespaces/{namespace}/ingresses`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Ingress](#ingress-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Ingress](#ingress-v1beta1) | OK


## Replace

> Examples using curl coming soon...

replace the specified Ingress

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/ingresses/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Ingress
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Ingress](#ingress-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Ingress](#ingress-v1beta1) | OK


## Patch

> Examples using curl coming soon...

partially update the specified Ingress

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/ingresses/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Ingress
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Ingress](#ingress-v1beta1) | OK


## Delete

> Examples using curl coming soon...

delete a Ingress

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/ingresses/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Ingress
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

read the specified Ingress

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/ingresses/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Ingress
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
200 | [Ingress](#ingress-v1beta1) | OK


## List

> Examples using curl coming soon...

list or watch objects of kind Ingress

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/ingresses`

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
200 | [IngressList](#ingresslist-v1beta1) | OK


## Watch

> Examples using curl coming soon...

watch changes to an object of kind Ingress

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/ingresses/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the Ingress
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

partially update status of the specified Ingress

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/ingresses/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Ingress
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Ingress](#ingress-v1beta1) | OK


## Read Status

> Examples using curl coming soon...

read status of the specified Ingress

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/ingresses/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Ingress
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Ingress](#ingress-v1beta1) | OK


## Replace Status

> Examples using curl coming soon...

replace status of the specified Ingress

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/ingresses/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Ingress
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Ingress](#ingress-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Ingress](#ingress-v1beta1) | OK


## Delete Collection

> Examples using curl coming soon...

delete collection of Ingress

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/ingresses`

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

list or watch objects of kind Ingress

### HTTP Request

`GET /apis/extensions/v1beta1/ingresses`

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
200 | [IngressList](#ingresslist-v1beta1) | OK


## Watch List

> Examples using curl coming soon...

watch individual changes to a list of Ingress

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/ingresses`

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

watch individual changes to a list of Ingress

### HTTP Request

`GET /apis/extensions/v1beta1/watch/ingresses`

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




