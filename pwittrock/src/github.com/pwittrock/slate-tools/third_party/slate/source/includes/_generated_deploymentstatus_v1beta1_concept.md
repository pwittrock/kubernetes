

-----------
# DeploymentStatus v1beta1

Group        | Version     | Kind
------------ | ---------- | -----------
Extensions | v1beta1 | DeploymentStatus







> Example yaml coming soon...


DeploymentStatus is the most recently observed status of the Deployment.

<aside class="notice">
Appears In <a href="#deployment-v1beta1">Deployment</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
availableReplicas | integer | Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
observedGeneration | integer | The generation observed by the deployment controller.
replicas | integer | Total number of non-terminated pods targeted by this deployment (their labels match the selector).
unavailableReplicas | integer | Total number of unavailable pods targeted by this deployment.
updatedReplicas | integer | Total number of non-terminated pods targeted by this deployment that have the desired template spec.





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



replace status of the specified Deployment

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Deployment
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Deployment](#deployment-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Deployment](#deployment-v1beta1) | OK


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



partially update status of the specified Deployment

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Deployment
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Deployment](#deployment-v1beta1) | OK



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



read status of the specified Deployment

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Deployment
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Deployment](#deployment-v1beta1) | OK




