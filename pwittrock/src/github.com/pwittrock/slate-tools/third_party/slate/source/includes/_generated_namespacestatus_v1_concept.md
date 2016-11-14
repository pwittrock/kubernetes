

-----------
# NamespaceStatus v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | NamespaceStatus







NamespaceStatus is information about the current status of a Namespace.

<aside class="notice">
Appears In <a href="#namespace-v1">Namespace</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
phase | string | Phase is the current lifecycle phase of the namespace. More info: http://releases.k8s.io/HEAD/docs/design/namespaces.md#phases





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



replace status of the specified Namespace

### HTTP Request

`PUT /api/v1/namespaces/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Namespace
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Namespace](#namespace-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Namespace](#namespace-v1) | OK


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



partially update status of the specified Namespace

### HTTP Request

`PATCH /api/v1/namespaces/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Namespace
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Namespace](#namespace-v1) | OK



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



read status of the specified Namespace

### HTTP Request

`GET /api/v1/namespaces/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Namespace
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Namespace](#namespace-v1) | OK




