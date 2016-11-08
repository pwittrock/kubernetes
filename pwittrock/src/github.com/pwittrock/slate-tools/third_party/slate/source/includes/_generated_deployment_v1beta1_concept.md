

-----------
# Deployment v1beta1

 > Deployment Config to run 3 nginx instances (max rollback set to 10 revisions). 

```shell
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: deployment-example
spec:
  replicas: 3
  revisionHistoryLimit: 10
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
kind: Deployment
metadata:
  name: deployment-example
spec:
  replicas: 3
  revisionHistoryLimit: 10
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
Extensions | v1beta1 | Deployment







Deployment enables declarative updates for Pods and ReplicaSets.

<aside class="notice">
Appears In <a href="#deploymentlist-v1beta1">DeploymentList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object metadata.
spec | [DeploymentSpec](#deploymentspec-v1beta1) | Specification of the desired behavior of the Deployment.
status | [DeploymentStatus](#deploymentstatus-v1beta1) | Most recently observed status of the Deployment.


### DeploymentSpec v1beta1

<aside class="notice">
Appears In <a href="#deployment-v1beta1">Deployment</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
minReadySeconds | integer | Minimum number of seconds for which a newly created pod should be ready without any of its container crashing, for it to be considered available. Defaults to 0 (pod will be considered available as soon as it is ready)
paused | boolean | Indicates that the deployment is paused and will not be processed by the deployment controller.
replicas | integer | Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
revisionHistoryLimit | integer | The number of old ReplicaSets to retain to allow rollback. This is a pointer to distinguish between explicit zero and not specified.
rollbackTo | [RollbackConfig](#rollbackconfig-v1beta1) | The config this deployment is rolling back to. Will be cleared after rollback is done.
selector | [LabelSelector](#labelselector-v1beta1) | Label selector for pods. Existing ReplicaSets whose pods are selected by this will be the ones affected by this deployment.
strategy | [DeploymentStrategy](#deploymentstrategy-v1beta1) | The deployment strategy to use to replace existing pods with new ones.
template | [PodTemplateSpec](#podtemplatespec-v1) | Template describes the pods that will be created.

### DeploymentStatus v1beta1

<aside class="notice">
Appears In <a href="#deployment-v1beta1">Deployment</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
availableReplicas | integer | Total number of available pods (ready for at least minReadySeconds) targeted by this deployment.
observedGeneration | integer | The generation observed by the deployment controller.
replicas | integer | Total number of non-terminated pods targeted by this deployment (their labels match the selector).
unavailableReplicas | integer | Total number of unavailable pods targeted by this deployment.
updatedReplicas | integer | Total number of non-terminated pods targeted by this deployment that have the desired template spec.

### DeploymentList v1beta1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [Deployment](#deployment-v1beta1) array | Items is the list of Deployments.
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata.

### DeploymentStrategy v1beta1

<aside class="notice">
Appears In <a href="#deploymentspec-v1beta1">DeploymentSpec</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
rollingUpdate | [RollingUpdateDeployment](#rollingupdatedeployment-v1beta1) | Rolling update config params. Present only if DeploymentStrategyType = RollingUpdate.
type | string | Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.

### DeploymentRollback v1beta1



Field        | Schema     | Description
------------ | ---------- | -----------
name | string | Required: This must match the Name of a deployment.
rollbackTo | [RollbackConfig](#rollbackconfig-v1beta1) | The config of this deployment rollback.
updatedAnnotations | object | The annotations to be updated to a deployment

### RollingUpdateDeployment v1beta1

<aside class="notice">
Appears In <a href="#deploymentstrategy-v1beta1">DeploymentStrategy</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
maxSurge | [IntOrString](#intorstring-intstr) | The maximum number of pods that can be scheduled above the desired number of pods. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). This can not be 0 if MaxUnavailable is 0. Absolute number is calculated from percentage by rounding up. By default, a value of 1 is used. Example: when this is set to 30%, the new RC can be scaled up immediately when the rolling update starts, such that the total number of old and new pods do not exceed 130% of desired pods. Once old pods have been killed, new RC can be scaled up further, ensuring that total number of pods running at any time during the update is atmost 130% of desired pods.
maxUnavailable | [IntOrString](#intorstring-intstr) | The maximum number of pods that can be unavailable during the update. Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%). Absolute number is calculated from percentage by rounding up. This can not be 0 if MaxSurge is 0. By default, a fixed value of 1 is used. Example: when this is set to 30%, the old RC can be scaled down to 70% of desired pods immediately when the rolling update starts. Once new pods are ready, old RC can be scaled down further, followed by scaling up the new RC, ensuring that the total number of pods available at all times during the update is at least 70% of desired pods.




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Execute

```shell

$ echo 'apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: deployment-example
spec:
  replicas: 3
  revisionHistoryLimit: 10
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.10
        ports:
        - containerPort: 80
' | kubectl create -f -

```



```yaml

$ kubectl proxy
$ curl -X POST -H 'Content-Type: application/yaml' --data '
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: deployment-example
spec:
  replicas: 3
  revisionHistoryLimit: 10
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.10
        ports:
        - containerPort: 80
' http://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/default/deployments

```

> Returns

```shell

deployment "deployment-example" created

```


```yaml

{
  "kind": "Deployment",
  "apiVersion": "extensions/v1beta1",
  "metadata": {
    "name": "deployment-example",
    "namespace": "default",
    "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example",
    "uid": "4ccca349-9cb1-11e6-9c54-42010a800148",
    "resourceVersion": "2118306",
    "generation": 1,
    "creationTimestamp": "2016-10-28T01:53:19Z",
    "labels": {
      "app": "nginx"
    }
  },
  "spec": {
    "replicas": 3,
    "selector": {
      "matchLabels": {
        "app": "nginx"
      }
    },
    "template": {
      "metadata": {
        "creationTimestamp": null,
        "labels": {
          "app": "nginx"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "nginx",
            "image": "nginx:1.10",
            "ports": [
              {
                "containerPort": 80,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "terminationMessagePath": "/dev/termination-log",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "securityContext": {}
      }
    },
    "strategy": {
      "type": "RollingUpdate",
      "rollingUpdate": {
        "maxUnavailable": 1,
        "maxSurge": 1
      }
    },
    "revisionHistoryLimit": 10
  },
  "status": {}
}

```



create a Deployment

### HTTP Request

`POST /apis/extensions/v1beta1/namespaces/{namespace}/deployments`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
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


## Replace

> Execute

```shell

$ echo 'apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: deployment-example
spec:
  replicas: 3
  revisionHistoryLimit: 10
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.11
        ports:
        - containerPort: 80
' | kubectl replace -f -

```



```yaml

$ kubectl proxy
$ curl -X PUT -H 'Content-Type: application/yaml' --data '
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: deployment-example
spec:
  replicas: 3
  revisionHistoryLimit: 10
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.11
        ports:
        - containerPort: 80
' http://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example

```

> Returns

```shell

deployment "deployment-example" replaced

```


```yaml

{
  "kind": "Deployment",
  "apiVersion": "extensions/v1beta1",
  "metadata": {
    "name": "deployment-example",
    "namespace": "default",
    "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example",
    "uid": "4ccca349-9cb1-11e6-9c54-42010a800148",
    "resourceVersion": "2119082",
    "generation": 5,
    "creationTimestamp": "2016-10-28T01:53:19Z",
    "labels": {
      "app": "nginx"
    }
  },
  "spec": {
    "replicas": 3,
    "selector": {
      "matchLabels": {
        "app": "nginx"
      }
    },
    "template": {
      "metadata": {
        "creationTimestamp": null,
        "labels": {
          "app": "nginx"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "nginx",
            "image": "nginx:1.11",
            "ports": [
              {
                "containerPort": 80,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "terminationMessagePath": "/dev/termination-log",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "securityContext": {}
      }
    },
    "strategy": {
      "type": "RollingUpdate",
      "rollingUpdate": {
        "maxUnavailable": 1,
        "maxSurge": 1
      }
    },
    "revisionHistoryLimit": 10
  },
  "status": {
    "observedGeneration": 4,
    "replicas": 3,
    "updatedReplicas": 3,
    "availableReplicas": 3
  }
}


```



replace the specified Deployment

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}`

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

$ kubectl patch deployment deployment-example -p \
	'{"spec":{"containers":[{"name":"nginx","image":"nginx:1.11"}]}}'

```



```yaml

$ kubectl proxy
$ curl -X PATCH -H 'Content-Type: application/strategic-merge-patch+json' --data '
{"spec":{"containers":[{"name":"nginx","image":"nginx:1.11"}]}}' \
	'http://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example'

```

> Returns

```shell

"deployment-example" patched

```


```yaml

{
  "kind": "Deployment",
  "apiVersion": "extensions/v1beta1",
  "metadata": {
    "name": "deployment-example",
    "namespace": "default",
    "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example",
    "uid": "d3be4f35-9ccc-11e6-9c54-42010a800148",
    "resourceVersion": "2137435",
    "generation": 2,
    "creationTimestamp": "2016-10-28T05:10:22Z",
    "labels": {
      "app": "nginx"
    },
    "annotations": {
      "deployment.kubernetes.io/revision": "3"
    }
  },
  "spec": {
    "replicas": 3,
    "selector": {
      "matchLabels": {
        "app": "nginx"
      }
    },
    "template": {
      "metadata": {
        "creationTimestamp": null,
        "labels": {
          "app": "nginx"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "nginx",
            "image": "nginx:1.10",
            "ports": [
              {
                "containerPort": 80,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "terminationMessagePath": "/dev/termination-log",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "securityContext": {}
      }
    },
    "strategy": {
      "type": "RollingUpdate",
      "rollingUpdate": {
        "maxUnavailable": 1,
        "maxSurge": 1
      }
    },
    "revisionHistoryLimit": 10
  },
  "status": {
    "observedGeneration": 2,
    "replicas": 3,
    "updatedReplicas": 3,
    "availableReplicas": 3
  }
}

```



partially update the specified Deployment

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}`

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


## Delete

> Execute

```shell

$ kubectl delete deployment deployment-example

```



```yaml

$ kubectl proxy
$ curl -X DELETE -H 'Content-Type: application/yaml' --data '
gracePeriodSeconds: 0
orphanDependents: false
' 'http://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example'

```

> Returns

```shell

deployment "deployment-example" deleted

```


```yaml

{
  "kind": "Status",
  "apiVersion": "v1",
  "metadata": {},
  "status": "Success",
  "code": 200
}


```



delete a Deployment

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Deployment
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

$ kubectl get deployment deployment-example -o json

```



```yaml

$ kubectl proxy
$ curl -X GET http://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example

```

> Returns

```shell

{
  "kind": "Deployment",
  "apiVersion": "extensions/v1beta1",
  "metadata": {
    "name": "deployment-example",
    "namespace": "default",
    "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example",
    "uid": "1b33145a-9c63-11e6-9c54-42010a800148",
    "resourceVersion": "2064726",
    "generation": 4,
    "creationTimestamp": "2016-10-27T16:33:35Z",
    "labels": {
      "app": "nginx"
    },
    "annotations": {
      "deployment.kubernetes.io/revision": "1"
    }
  },
  "spec": {
    "replicas": 3,
    "selector": {
      "matchLabels": {
        "app": "nginx"
      }
    },
    "template": {
      "metadata": {
        "creationTimestamp": null,
        "labels": {
          "app": "nginx"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "nginx",
            "image": "nginx:1.10",
            "ports": [
              {
                "containerPort": 80,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "terminationMessagePath": "/dev/termination-log",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "securityContext": {}
      }
    },
    "strategy": {
      "type": "RollingUpdate",
      "rollingUpdate": {
        "maxUnavailable": 1,
        "maxSurge": 1
      }
    }
  },
  "status": {
    "observedGeneration": 4,
    "replicas": 3,
    "updatedReplicas": 3,
    "availableReplicas": 3
  }
}


```


```yaml

{
  "kind": "Deployment",
  "apiVersion": "extensions/v1beta1",
  "metadata": {
    "name": "deployment-example",
    "namespace": "default",
    "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example",
    "uid": "1b33145a-9c63-11e6-9c54-42010a800148",
    "resourceVersion": "2064726",
    "generation": 4,
    "creationTimestamp": "2016-10-27T16:33:35Z",
    "labels": {
      "app": "nginx"
    },
    "annotations": {
      "deployment.kubernetes.io/revision": "1"
    }
  },
  "spec": {
    "replicas": 3,
    "selector": {
      "matchLabels": {
        "app": "nginx"
      }
    },
    "template": {
      "metadata": {
        "creationTimestamp": null,
        "labels": {
          "app": "nginx"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "nginx",
            "image": "nginx:1.10",
            "ports": [
              {
                "containerPort": 80,
                "protocol": "TCP"
              }
            ],
            "resources": {},
            "terminationMessagePath": "/dev/termination-log",
            "imagePullPolicy": "IfNotPresent"
          }
        ],
        "restartPolicy": "Always",
        "terminationGracePeriodSeconds": 30,
        "dnsPolicy": "ClusterFirst",
        "securityContext": {}
      }
    },
    "strategy": {
      "type": "RollingUpdate",
      "rollingUpdate": {
        "maxUnavailable": 1,
        "maxSurge": 1
      }
    }
  },
  "status": {
    "observedGeneration": 4,
    "replicas": 3,
    "updatedReplicas": 3,
    "availableReplicas": 3
  }
}


```



read the specified Deployment

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Deployment
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
200 | [Deployment](#deployment-v1beta1) | OK


## List

> Execute

```shell

$ kubectl get deployment -o json

```



```yaml

$ kubectl proxy
$ curl -X GET 'http://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/default/deployments'

```

> Returns

```shell

{
  "kind": "List",
  "apiVersion": "v1",
  "metadata": {},
  "items": [
    {
      "kind": "Deployment",
      "apiVersion": "extensions/v1beta1",
      "metadata": {
        "name": "docs",
        "namespace": "default",
        "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/docs",
        "uid": "ef49e1d2-915e-11e6-be81-42010a80003f",
        "resourceVersion": "1924126",
        "generation": 21,
        "creationTimestamp": "2016-10-13T16:06:00Z",
        "labels": {
          "run": "docs"
        },
        "annotations": {
          "deployment.kubernetes.io/revision": "10",
          "replicatingperfection.net/push-image": "true"
        }
      },
      "spec": {
        "replicas": 1,
        "selector": {
          "matchLabels": {
            "run": "docs"
          }
        },
        "template": {
          "metadata": {
            "creationTimestamp": null,
            "labels": {
              "auto-pushed-image-pwittrock/api-docs": "1477496453",
              "run": "docs"
            }
          },
          "spec": {
            "containers": [
              {
                "name": "docs",
                "image": "pwittrock/api-docs:v9",
                "resources": {},
                "terminationMessagePath": "/dev/termination-log",
                "imagePullPolicy": "Always"
              }
            ],
            "restartPolicy": "Always",
            "terminationGracePeriodSeconds": 30,
            "dnsPolicy": "ClusterFirst",
            "securityContext": {}
          }
        },
        "strategy": {
          "type": "RollingUpdate",
          "rollingUpdate": {
            "maxUnavailable": 1,
            "maxSurge": 1
          }
        }
      },
      "status": {
        "observedGeneration": 21,
        "replicas": 1,
        "updatedReplicas": 1,
        "availableReplicas": 1
      }
    },
    {
      "kind": "Deployment",
      "apiVersion": "extensions/v1beta1",
      "metadata": {
        "name": "deployment-example",
        "namespace": "default",
        "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example",
        "uid": "1b33145a-9c63-11e6-9c54-42010a800148",
        "resourceVersion": "2064726",
        "generation": 4,
        "creationTimestamp": "2016-10-27T16:33:35Z",
        "labels": {
          "app": "nginx"
        },
        "annotations": {
          "deployment.kubernetes.io/revision": "1"
        }
      },
      "spec": {
        "replicas": 3,
        "selector": {
          "matchLabels": {
            "app": "nginx"
          }
        },
        "template": {
          "metadata": {
            "creationTimestamp": null,
            "labels": {
              "app": "nginx"
            }
          },
          "spec": {
            "containers": [
              {
                "name": "nginx",
                "image": "nginx:1.10",
                "ports": [
                  {
                    "containerPort": 80,
                    "protocol": "TCP"
                  }
                ],
                "resources": {},
                "terminationMessagePath": "/dev/termination-log",
                "imagePullPolicy": "IfNotPresent"
              }
            ],
            "restartPolicy": "Always",
            "terminationGracePeriodSeconds": 30,
            "dnsPolicy": "ClusterFirst",
            "securityContext": {}
          }
        },
        "strategy": {
          "type": "RollingUpdate",
          "rollingUpdate": {
            "maxUnavailable": 1,
            "maxSurge": 1
          }
        }
      },
      "status": {
        "observedGeneration": 4,
        "replicas": 3,
        "updatedReplicas": 3,
        "availableReplicas": 3
      }
    }
  ]
}


```


```yaml

{
  "kind": "List",
  "apiVersion": "v1",
  "metadata": {},
  "items": [
    {
      "kind": "Deployment",
      "apiVersion": "extensions/v1beta1",
      "metadata": {
        "name": "docs",
        "namespace": "default",
        "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/docs",
        "uid": "ef49e1d2-915e-11e6-be81-42010a80003f",
        "resourceVersion": "1924126",
        "generation": 21,
        "creationTimestamp": "2016-10-13T16:06:00Z",
        "labels": {
          "run": "docs"
        },
        "annotations": {
          "deployment.kubernetes.io/revision": "10",
          "replicatingperfection.net/push-image": "true"
        }
      },
      "spec": {
        "replicas": 1,
        "selector": {
          "matchLabels": {
            "run": "docs"
          }
        },
        "template": {
          "metadata": {
            "creationTimestamp": null,
            "labels": {
              "auto-pushed-image-pwittrock/api-docs": "1477496453",
              "run": "docs"
            }
          },
          "spec": {
            "containers": [
              {
                "name": "docs",
                "image": "pwittrock/api-docs:v9",
                "resources": {},
                "terminationMessagePath": "/dev/termination-log",
                "imagePullPolicy": "Always"
              }
            ],
            "restartPolicy": "Always",
            "terminationGracePeriodSeconds": 30,
            "dnsPolicy": "ClusterFirst",
            "securityContext": {}
          }
        },
        "strategy": {
          "type": "RollingUpdate",
          "rollingUpdate": {
            "maxUnavailable": 1,
            "maxSurge": 1
          }
        }
      },
      "status": {
        "observedGeneration": 21,
        "replicas": 1,
        "updatedReplicas": 1,
        "availableReplicas": 1
      }
    },
    {
      "kind": "Deployment",
      "apiVersion": "extensions/v1beta1",
      "metadata": {
        "name": "deployment-example",
        "namespace": "default",
        "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example",
        "uid": "1b33145a-9c63-11e6-9c54-42010a800148",
        "resourceVersion": "2064726",
        "generation": 4,
        "creationTimestamp": "2016-10-27T16:33:35Z",
        "labels": {
          "app": "nginx"
        },
        "annotations": {
          "deployment.kubernetes.io/revision": "1"
        }
      },
      "spec": {
        "replicas": 3,
        "selector": {
          "matchLabels": {
            "app": "nginx"
          }
        },
        "template": {
          "metadata": {
            "creationTimestamp": null,
            "labels": {
              "app": "nginx"
            }
          },
          "spec": {
            "containers": [
              {
                "name": "nginx",
                "image": "nginx:1.10",
                "ports": [
                  {
                    "containerPort": 80,
                    "protocol": "TCP"
                  }
                ],
                "resources": {},
                "terminationMessagePath": "/dev/termination-log",
                "imagePullPolicy": "IfNotPresent"
              }
            ],
            "restartPolicy": "Always",
            "terminationGracePeriodSeconds": 30,
            "dnsPolicy": "ClusterFirst",
            "securityContext": {}
          }
        },
        "strategy": {
          "type": "RollingUpdate",
          "rollingUpdate": {
            "maxUnavailable": 1,
            "maxSurge": 1
          }
        }
      },
      "status": {
        "observedGeneration": 4,
        "replicas": 3,
        "updatedReplicas": 3,
        "availableReplicas": 3
      }
    }
  ]
}


```



list or watch objects of kind Deployment

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/deployments`

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
200 | [DeploymentList](#deploymentlist-v1beta1) | OK


## Watch

> Execute

```shell

$ kubectl get deployment deployment-example --watch -o json

```



```yaml

$ kubectl proxy
$ curl -X GET 'http://127.0.0.1:8001/apis/extensions/v1beta1/watch/namespaces/default/deployments/deployment-example'

```

> Returns

```shell

{
	"type": "ADDED",
	"object": {
		"kind": "Deployment",
		"apiVersion": "extensions/v1beta1",
		"metadata": {
			"name": "deployment-example",
			"namespace": "default",
			"selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example",
			"uid": "64c12290-9cbf-11e6-9c54-42010a800148",
			"resourceVersion": "2128095",
			"generation": 2,
			"creationTimestamp": "2016-10-28T03:34:12Z",
			"labels": {
				"app": "nginx"
			},
			"annotations": {
				"deployment.kubernetes.io/revision": "3"
			}
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"creationTimestamp": null,
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [
						{
							"name": "nginx",
							"image": "nginx:1.10",
							"ports": [
								{
									"containerPort": 80,
									"protocol": "TCP"
								}
							],
							"resources": {
							},
							"terminationMessagePath": "/dev/termination-log",
							"imagePullPolicy": "IfNotPresent"
						}
					],
					"restartPolicy": "Always",
					"terminationGracePeriodSeconds": 30,
					"dnsPolicy": "ClusterFirst",
					"securityContext": {
					}
				}
			},
			"strategy": {
				"type": "RollingUpdate",
				"rollingUpdate": {
					"maxUnavailable": 1,
					"maxSurge": 1
				}
			}
		},
		"status": {
			"observedGeneration": 2,
			"replicas": 3,
			"updatedReplicas": 3,
			"availableReplicas": 3
		}
	}
}

```


```yaml

{
	"type": "ADDED",
	"object": {
		"kind": "Deployment",
		"apiVersion": "extensions/v1beta1",
		"metadata": {
			"name": "deployment-example",
			"namespace": "default",
			"selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/deployment-example",
			"uid": "64c12290-9cbf-11e6-9c54-42010a800148",
			"resourceVersion": "2128095",
			"generation": 2,
			"creationTimestamp": "2016-10-28T03:34:12Z",
			"labels": {
				"app": "nginx"
			},
			"annotations": {
				"deployment.kubernetes.io/revision": "3"
			}
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"creationTimestamp": null,
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [
						{
							"name": "nginx",
							"image": "nginx:1.10",
							"ports": [
								{
									"containerPort": 80,
									"protocol": "TCP"
								}
							],
							"resources": {
							},
							"terminationMessagePath": "/dev/termination-log",
							"imagePullPolicy": "IfNotPresent"
						}
					],
					"restartPolicy": "Always",
					"terminationGracePeriodSeconds": 30,
					"dnsPolicy": "ClusterFirst",
					"securityContext": {
					}
				}
			},
			"strategy": {
				"type": "RollingUpdate",
				"rollingUpdate": {
					"maxUnavailable": 1,
					"maxSurge": 1
				}
			}
		},
		"status": {
			"observedGeneration": 2,
			"replicas": 3,
			"updatedReplicas": 3,
			"availableReplicas": 3
		}
	}
}

```



watch changes to an object of kind Deployment

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/deployments/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the Deployment
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



delete collection of Deployment

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/deployments`

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



list or watch objects of kind Deployment

### HTTP Request

`GET /apis/extensions/v1beta1/deployments`

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
200 | [DeploymentList](#deploymentlist-v1beta1) | OK


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



watch individual changes to a list of Deployment

### HTTP Request

`GET /apis/extensions/v1beta1/watch/deployments`

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

`GET /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}/scale`

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

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}/scale`

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

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}/scale`

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


## Rollback

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



create rollback of a DeploymentRollback

### HTTP Request

`POST /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}/rollback`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [DeploymentRollback](#deploymentrollback-v1beta1) | 
name |  | name of the DeploymentRollback
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [DeploymentRollback](#deploymentrollback-v1beta1) | OK




