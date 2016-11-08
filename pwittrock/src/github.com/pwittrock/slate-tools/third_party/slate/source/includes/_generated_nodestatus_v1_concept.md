

-----------
# NodeStatus v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | NodeStatus







NodeStatus is information about the current status of a node.

<aside class="notice">
Appears In <a href="#node-v1">Node</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
addresses | [NodeAddress](#nodeaddress-v1) array | List of addresses reachable to the node. Queried from cloud provider, if available. More info: http://releases.k8s.io/HEAD/docs/admin/node.md#node-addresses
allocatable | object | Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
capacity | object | Capacity represents the total resources of a node. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#capacity for more details.
conditions | [NodeCondition](#nodecondition-v1) array | Conditions is an array of current observed node conditions. More info: http://releases.k8s.io/HEAD/docs/admin/node.md#node-condition
daemonEndpoints | [NodeDaemonEndpoints](#nodedaemonendpoints-v1) | Endpoints of daemons running on the Node.
images | [ContainerImage](#containerimage-v1) array | List of container images on this node
nodeInfo | [NodeSystemInfo](#nodesysteminfo-v1) | Set of ids/uuids to uniquely identify the node. More info: http://releases.k8s.io/HEAD/docs/admin/node.md#node-info
phase | string | NodePhase is the recently observed lifecycle phase of the node. More info: http://releases.k8s.io/HEAD/docs/admin/node.md#node-phase The field is never populated, and now is deprecated.
volumesAttached | [AttachedVolume](#attachedvolume-v1) array | List of volumes that are attached to the node.
volumesInUse | string array | List of attachable volumes in use (mounted) by the node.





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



replace status of the specified Node

### HTTP Request

`PUT /api/v1/nodes/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Node
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Node](#node-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Node](#node-v1) | OK


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



partially update status of the specified Node

### HTTP Request

`PATCH /api/v1/nodes/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Node
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Node](#node-v1) | OK



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



read status of the specified Node

### HTTP Request

`GET /api/v1/nodes/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Node
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Node](#node-v1) | OK




