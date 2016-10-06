# Deployment

Field                         | Operations     | Description
----------------------------- | -------------- | -----------
apiVersion                    | all            | Must be set to 'extensions/v1beta1'
kind                          | all            | Must be set to 'Deployment'
metadata.name                 | create, update | Unique name for the object
spec.replicas                 | create, update | Number of desired pods. This is a pointer to distinguish between explicit zero and not specified. Defaults to 1.
spec.template.metadata.labels | create, update | 
spec.template.spec            | create, update | Specification of the desired behavior of the pod.

## Get Deployments

> Run `$ kubectl proxy`

```shell
curl 'http://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/default/deployments?labelSelector=app=nginx'
```

```ruby
TODO: Write this
```

> The above command returns JSON structured like this:

```shell
{
  "kind": "DeploymentList",
  "apiVersion": "extensions/v1beta1",
  "metadata": {
    "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments",
    "resourceVersion": "209615"
  },
  "items": [
    {
      "metadata": {
        "name": "nginx-deployment",
        "namespace": "default",
        "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/nginx-deployment",
        "uid": "0504efce-8ca9-11e6-ae2a-42010a80003f",
        "resourceVersion": "209590",
        "generation": 4,
        "creationTimestamp": "2016-10-07T16:13:44Z",
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
                "image": "nginx",
                "ports": [
                  {
                    "containerPort": 80,
                    "protocol": "TCP"
                  }
                ],
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
        "observedGeneration": 4,
        "replicas": 3,
        "updatedReplicas": 3,
        "availableReplicas": 3
      }
    }
  ]
}
```

```ruby
TODO: Write this
```

List or watch objects of kind Deployment.

### HTTP Request

`GET http://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/{namespace}/deployments`

### Path Parameters

Parameter | Description
--------- | -----------
namespace | object name and auth scope, such as for teams and projects

### Query Parameters

Parameter | Default | Description | Required | Allow Multiple
--------- | ------- | ----------- | -------- | --------------
pretty    | false   | If 'true', then the output is pretty printed. | false | false
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything. | false | false
fieldSelector | false | A selector to restrict the list of returned objects by their fields. Defaults to everything. | false | false
watch | false | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion. | false | false
resourceVersion | | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history. | false  | false
timeoutSeconds | | Timeout for the list/watch call. | false | false

## Create a Deployment

> Run `$ kubectl proxy`

```shell
curl -H 'Content-Type: application/yaml' --data '
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx
        ports:
        - containerPort: 80
' 'http://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/default/deployments'
```

> The above command returns JSON structured like this:

```json
{
  "kind": "Deployment",
  "apiVersion": "extensions/v1beta1",
  "metadata": {
    "name": "nginx-deployment",
    "namespace": "default",
    "selfLink": "/apis/extensions/v1beta1/namespaces/default/deployments/nginx-deployment",
    "uid": "0504efce-8ca9-11e6-ae2a-42010a80003f",
    "resourceVersion": "208985",
    "generation": 1,
    "creationTimestamp": "2016-10-07T16:13:44Z",
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
            "image": "nginx",
            "ports": [
              {
                "containerPort": 80,
                "protocol": "TCP"
              }
            ],
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
  "status": {}
}
```

### HTTP Request

`POST https://127.0.0.1:8001/apis/extensions/v1beta1/namespaces/{namespace}/deployments`

### Path Parameters

Parameter | Schema | Description
--------- | ------ | -----------
namespace | string | Object name and auth scope, such as for teams and projects

### Body Parameters

Parameter     | Schema                    | Description
------------- | ------------------------- | -----------
body          | [Deployment](#deployment) | 

## Replace a Deployment



<aside class="notice">
It is recommended to use <a href="#patch-a-deployment">PATCH</a>
instead of Replace in order to retain fields set by the server.
</aside>

> Run `$ kubectl proxy`

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}`

### Path Parameters

Parameter | Description
--------- | -----------
namespace | object name and auth scope, such as for teams and projects
name      | name of the Deployment

### Query Parameters

Parameter | Description
--------- | -----------
pretty    | If true, then the output is pretty printed.

## Patch a Deployment

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}`

### Path Parameters

Parameter | Description
--------- | -----------
namespace | object name and auth scope, such as for teams and projects
name      | name of the Deployment

### Query Parameters

Parameter | Description
--------- | -----------
pretty    | If true, then the output is pretty printed.

## Delete a Deployment

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/deployments/{name}`

### Path Parameters

Parameter | Description
--------- | -----------
namespace | object name and auth scope, such as for teams and projects
name      | name of the Deployment

### Query Parameters

Parameter | Description
--------- | -----------
pretty    | If true, then the output is pretty printed.

