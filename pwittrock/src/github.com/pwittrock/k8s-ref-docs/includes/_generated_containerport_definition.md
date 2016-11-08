## *ContainerPort v1*

> Example yaml coming soon...



ContainerPort represents a network port in a single container.

<aside class="notice">
Appears In  <a href="#container-v1">Container</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
containerPort | integer | Number of port to expose on the pod's IP address. This must be a valid port number, 0 < x < 65536.
hostIP | string | What host IP to bind the external port to.
hostPort | integer | Number of port to expose on the host. If specified, this must be a valid port number, 0 < x < 65536. If HostNetwork is specified, this must match ContainerPort. Most containers do not need this.
name | string | If specified, this must be an IANA_SVC_NAME and unique within the pod. Each named port in a pod must have a unique name. Name for the port that can be referred to by services.
protocol | string | Protocol for port. Must be UDP or TCP. Defaults to "TCP".

