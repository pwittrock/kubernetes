

-----------
# ResourceQuotaStatus v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ResourceQuotaStatus







ResourceQuotaStatus defines the enforced hard limits and observed use.

<aside class="notice">
Appears In <a href="#resourcequota-v1">ResourceQuota</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
hard | object | Hard is the set of enforced hard limits for each named resource. More info: http://releases.k8s.io/HEAD/docs/design/admission_control_resource_quota.md#admissioncontrol-plugin-resourcequota
used | object | Used is the current observed total usage of the resource in the namespace.





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




