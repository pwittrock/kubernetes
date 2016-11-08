

-----------
# DaemonSetStatus v1beta1



Group        | Version     | Kind
------------ | ---------- | -----------
Extensions | v1beta1 | DaemonSetStatus







DaemonSetStatus represents the current status of a daemon set.

<aside class="notice">
Appears In <a href="#daemonset-v1beta1">DaemonSet</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
currentNumberScheduled | integer | CurrentNumberScheduled is the number of nodes that are running at least 1 daemon pod and are supposed to run the daemon pod. More info: http://releases.k8s.io/HEAD/docs/admin/daemons.md
desiredNumberScheduled | integer | DesiredNumberScheduled is the total number of nodes that should be running the daemon pod (including nodes correctly running the daemon pod). More info: http://releases.k8s.io/HEAD/docs/admin/daemons.md
numberMisscheduled | integer | NumberMisscheduled is the number of nodes that are running the daemon pod, but are not supposed to run the daemon pod. More info: http://releases.k8s.io/HEAD/docs/admin/daemons.md
numberReady | integer | NumberReady is the number of nodes that should be running the daemon pod and have one or more of the daemon pod running and ready.





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



replace status of the specified DaemonSet

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the DaemonSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [DaemonSet](#daemonset-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSet](#daemonset-v1beta1) | OK


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



partially update status of the specified DaemonSet

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the DaemonSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSet](#daemonset-v1beta1) | OK



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



read status of the specified DaemonSet

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/daemonsets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the DaemonSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DaemonSet](#daemonset-v1beta1) | OK




