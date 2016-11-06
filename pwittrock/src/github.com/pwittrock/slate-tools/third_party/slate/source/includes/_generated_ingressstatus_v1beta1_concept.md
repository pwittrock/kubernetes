

-----------
# IngressStatus v1beta1

Group        | Version     | Kind
------------ | ---------- | -----------
Extensions | v1beta1 | IngressStatus







> Example yaml coming soon...


IngressStatus describe the current state of the Ingress.

<aside class="notice">
Appears In <a href="#ingress-v1beta1">Ingress</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
loadBalancer | [LoadBalancerStatus](#loadbalancerstatus-v1) | LoadBalancer contains the current status of the load-balancer.





## <strong>Write Operations</strong>

See supported operations below...

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




