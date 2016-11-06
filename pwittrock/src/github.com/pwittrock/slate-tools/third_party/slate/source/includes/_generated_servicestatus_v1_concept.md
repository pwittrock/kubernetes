

-----------
# ServiceStatus v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ServiceStatus







> Example yaml coming soon...


ServiceStatus represents the current status of a service.

<aside class="notice">
Appears In <a href="#service-v1">Service</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
loadBalancer | [LoadBalancerStatus](#loadbalancerstatus-v1) | LoadBalancer contains the current status of the load-balancer, if one is present.





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



replace status of the specified Service

### HTTP Request

`PUT /api/v1/namespaces/{namespace}/services/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Service
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Service](#service-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Service](#service-v1) | OK


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



partially update status of the specified Service

### HTTP Request

`PATCH /api/v1/namespaces/{namespace}/services/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Service
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Service](#service-v1) | OK



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



read status of the specified Service

### HTTP Request

`GET /api/v1/namespaces/{namespace}/services/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Service
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Service](#service-v1) | OK




