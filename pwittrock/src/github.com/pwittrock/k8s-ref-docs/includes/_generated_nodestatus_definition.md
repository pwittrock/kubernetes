## *NodeStatus v1*

> Example yaml coming soon...



NodeStatus is information about the current status of a node.

<aside class="notice">
Appears In  <a href="#node-v1">Node</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
allocatable | object | Allocatable represents the resources of a node that are available for scheduling. Defaults to Capacity.
capacity | object | Capacity represents the total resources of a node. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#capacity for more details.
conditions | [NodeCondition](#nodecondition-v1) array | Conditions is an array of current observed node conditions. More info: http://releases.k8s.io/HEAD/docs/admin/node.md#node-condition
volumesAttached | [AttachedVolume](#attachedvolume-v1) array | List of volumes that are attached to the node.
volumesInUse | string array | List of attachable volumes in use (mounted) by the node.
addresses | [NodeAddress](#nodeaddress-v1) array | List of addresses reachable to the node. Queried from cloud provider, if available. More info: http://releases.k8s.io/HEAD/docs/admin/node.md#node-addresses
daemonEndpoints | [NodeDaemonEndpoints](#nodedaemonendpoints-v1) | Endpoints of daemons running on the Node.
images | [ContainerImage](#containerimage-v1) array | List of container images on this node
nodeInfo | [NodeSystemInfo](#nodesysteminfo-v1) | Set of ids/uuids to uniquely identify the node. More info: http://releases.k8s.io/HEAD/docs/admin/node.md#node-info
phase | string | NodePhase is the recently observed lifecycle phase of the node. More info: http://releases.k8s.io/HEAD/docs/admin/node.md#node-phase The field is never populated, and now is deprecated.

