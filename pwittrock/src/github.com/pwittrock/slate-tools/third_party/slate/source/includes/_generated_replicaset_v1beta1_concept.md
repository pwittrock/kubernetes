

-----------
# ReplicaSet v1beta1

 > ReplicaSet Config to run 3 nginx instances. 

```shell
apiVersion: extensions/v1beta1
kind: ReplicaSet
metadata:
  name: replicaset-example
spec:
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.10

```


```yaml
apiVersion: extensions/v1beta1
kind: ReplicaSet
metadata:
  name: replicaset-example
spec:
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.10

```


Group        | Version     | Kind
------------ | ---------- | -----------
Extensions | v1beta1 | ReplicaSet

<aside class="warning">In most cases it is better to create a <a href="#deployment-v1beta1">Deployment</a> instead as it will create new ReplicaSets and manage rolling out changes to them.</aside>





ReplicaSet represents the configuration of a ReplicaSet.

<aside class="notice">
Appears In <a href="#replicasetlist-v1beta1">ReplicaSetList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | If the Labels of a ReplicaSet are empty, they are defaulted to be the same as the Pod(s) that the ReplicaSet manages. Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [ReplicaSetSpec](#replicasetspec-v1beta1) | Spec defines the specification of the desired behavior of the ReplicaSet. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status
status | [ReplicaSetStatus](#replicasetstatus-v1beta1) | Status is the most recently observed status of the ReplicaSet. This data may be out of date by some window of time. Populated by the system. Read-only. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status


### ReplicaSetSpec v1beta1

<aside class="notice">
Appears In <a href="#replicaset-v1beta1">ReplicaSet</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
minReadySeconds | integer | Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
replicas | integer | Replicas is the number of desired replicas. This is a pointer to distinguish between explicit zero and unspecified. Defaults to 1. More info: http://kubernetes.io/docs/user-guide/replication-controller#what-is-a-replication-controller
selector | [LabelSelector](#labelselector-v1beta1) | Selector is a label query over pods that should match the replica count. If the selector is empty, it is defaulted to the labels present on the pod template. Label keys and values that must match in order to be controlled by this replica set. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
template | [PodTemplateSpec](#podtemplatespec-v1) | Template is the object that describes the pod that will be created if insufficient replicas are detected. More info: http://kubernetes.io/docs/user-guide/replication-controller#pod-template

### ReplicaSetStatus v1beta1

<aside class="notice">
Appears In <a href="#replicaset-v1beta1">ReplicaSet</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
availableReplicas | integer | The number of available replicas (ready for at least minReadySeconds) for this replica set.
conditions | [ReplicaSetCondition](#replicasetcondition-v1beta1) array | Represents the latest available observations of a replica set's current state.
fullyLabeledReplicas | integer | The number of pods that have labels matching the labels of the pod template of the replicaset.
observedGeneration | integer | ObservedGeneration reflects the generation of the most recently observed ReplicaSet.
readyReplicas | integer | The number of ready replicas for this replica set.
replicas | integer | Replicas is the most recently oberved number of replicas. More info: http://kubernetes.io/docs/user-guide/replication-controller#what-is-a-replication-controller

### ReplicaSetList v1beta1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [ReplicaSet](#replicaset-v1beta1) array | List of ReplicaSets. More info: http://kubernetes.io/docs/user-guide/replication-controller
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds




## <strong>Write Operations</strong>

See supported operations below...

## Create

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



create a ReplicaSet

### HTTP Request

`POST /apis/extensions/v1beta1/namespaces/{namespace}/replicasets`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ReplicaSet](#replicaset-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ReplicaSet](#replicaset-v1beta1) | OK


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



replace the specified ReplicaSet

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/replicasets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ReplicaSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ReplicaSet](#replicaset-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ReplicaSet](#replicaset-v1beta1) | OK


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



partially update the specified ReplicaSet

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/replicasets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ReplicaSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ReplicaSet](#replicaset-v1beta1) | OK


## Delete

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



delete a ReplicaSet

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/replicasets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ReplicaSet
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



read the specified ReplicaSet

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/replicasets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ReplicaSet
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
200 | [ReplicaSet](#replicaset-v1beta1) | OK


## List

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



list or watch objects of kind ReplicaSet

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/replicasets`

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
200 | [ReplicaSetList](#replicasetlist-v1beta1) | OK


## Watch

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



watch changes to an object of kind ReplicaSet

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/replicasets/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the ReplicaSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-versioned) | OK



## <strong>Status & Collection Operations</strong>

See supported operations below...

## Patch Status

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



partially update status of the specified ReplicaSet

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/replicasets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ReplicaSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ReplicaSet](#replicaset-v1beta1) | OK


## Read Status

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



read status of the specified ReplicaSet

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/replicasets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ReplicaSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ReplicaSet](#replicaset-v1beta1) | OK


## Replace Status

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



replace status of the specified ReplicaSet

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/replicasets/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the ReplicaSet
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [ReplicaSet](#replicaset-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [ReplicaSet](#replicaset-v1beta1) | OK


## Delete Collection

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



delete collection of ReplicaSet

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/replicasets`

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



list or watch objects of kind ReplicaSet

### HTTP Request

`GET /apis/extensions/v1beta1/replicasets`

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
200 | [ReplicaSetList](#replicasetlist-v1beta1) | OK


## Watch List

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



watch individual changes to a list of ReplicaSet

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/replicasets`

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
200 | [Event](#event-versioned) | OK


## Watch List All Namespaces

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



watch individual changes to a list of ReplicaSet

### HTTP Request

`GET /apis/extensions/v1beta1/watch/replicasets`

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
200 | [Event](#event-versioned) | OK



## <strong>Misc Operations</strong>

See supported operations below...

## Read Scale

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



read scale of the specified Scale

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/replicasets/{name}/scale`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Scale
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Scale](#scale-v1beta1) | OK


## Replace Scale

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



replace scale of the specified Scale

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/replicasets/{name}/scale`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Scale
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Scale](#scale-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Scale](#scale-v1beta1) | OK


## Patch Scale

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



partially update scale of the specified Scale

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/replicasets/{name}/scale`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Scale
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Scale](#scale-v1beta1) | OK




