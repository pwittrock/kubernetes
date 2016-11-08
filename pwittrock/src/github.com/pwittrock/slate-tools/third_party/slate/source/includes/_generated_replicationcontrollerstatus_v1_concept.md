

-----------
# ReplicationControllerStatus v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ReplicationControllerStatus







ReplicationControllerStatus represents the current status of a replication controller.

<aside class="notice">
Appears In <a href="#replicationcontroller-v1">ReplicationController</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
availableReplicas | integer | The number of available replicas (ready for at least minReadySeconds) for this replication controller.
conditions | [ReplicationControllerCondition](#replicationcontrollercondition-v1) array | Represents the latest available observations of a replication controller's current state.
fullyLabeledReplicas | integer | The number of pods that have labels matching the labels of the pod template of the replication controller.
observedGeneration | integer | ObservedGeneration reflects the generation of the most recently observed replication controller.
readyReplicas | integer | The number of ready replicas for this replication controller.
replicas | integer | Replicas is the most recently oberved number of replicas. More info: http://kubernetes.io/docs/user-guide/replication-controller#what-is-a-replication-controller





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



replace status of the specified ReplicationController

### HTTP Request

`PUT /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ReplicationController
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ReplicationController](#replicationcontroller-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ReplicationController](#replicationcontroller-v1) | OK


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



partially update status of the specified ReplicationController

### HTTP Request

`PATCH /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ReplicationController
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ReplicationController](#replicationcontroller-v1) | OK



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



read status of the specified ReplicationController

### HTTP Request

`GET /api/v1/namespaces/{namespace}/replicationcontrollers/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ReplicationController
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ReplicationController](#replicationcontroller-v1) | OK




