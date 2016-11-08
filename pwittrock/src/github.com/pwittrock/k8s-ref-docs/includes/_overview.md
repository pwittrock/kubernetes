# API Overview

Welcome to the Kubernetes API! You can use our API to access Kubernetes API endpoints to read and write Kubernetes
resource objects.

## Resource Object Categories

This is a highlevel overview of the basic types of resources provide by the Kubernetes API and their primary functions.

### Workload Resources

*Run your containers*.

Workloads are responsible for managing and running your containers on the cluster.

### Discovery & LB Resources

*Access your containers*.

These are responsible for stitching your workloads together into an externally accessible Loadbalanced Service.

### Config & Storage Resources

*Inject initialization data into your containers*

These are responsible for injecting data into your applications and persisting data externally to your container.

### Cluster Resources

*Manage your cluster*.

These are responsible for defining configuration of the cluster itself, and are generally only used by cluster operators.

### Meta Resources

*Configure resource meta behavior*.

These are responsible for configuring behavior of your other Resources within the Cluster.

------------

## Resource Objects

Resource objects have 3 components:

### *ResourceSpec*

This is defined by the user and describes the desired state of system.  Fill this in when creating or updating an
object.

### *ResourceStatus*

This is filled in by the server and reports the current state of the system.  Only kubernetes components should fill
this in

### *Resource ObjectMeta*

This is metadata about the resource, such as its name, type, api version, annotations, and labels.  This contains
fields that maybe updated both by the end user and the system (e.g. annotations)

### *ResourceList*

This is used to return a list of resources from an api endpoint.

------------

## Resource Operations

Most resources provide the following Operations:

### Create

Create operations will create the resource in the storage backend.  After a resource is create the system will apply
the desired state.

### Update

Updates come in 2 flavors:

- Replace
- Patch

**Replace**

Replacing a resource object will update the resource by replacing the existing spec with the provided one.  For
read-then-write operations this is safe because an optimistic lock failure will occur if the resource was modified
between the read and write.  *Note*: The *Resource*Status will be ignored by the system and will not be updated.
To update the status, one must invoke the specific status update operation.

Note: Replacing a resource object may not result immediately in changes being propagated to downstream objects.  For instance
replacing a *ConfigMap* or *Secret* resource will not result in all *Pod*s seeing the changes unless the *Pod*s are
restarted out of band.

**Patch**

Patch will apply a change to a specific field.  How the change is merged is defined per field.  Lists may either be
replaced or merged.  Merging lists will not preserve ordering.

**Patches will never cause optimistic locking failures, and the last write will win.**  Patches are recommended
 when the full state is not read before an update, or when failing on optimistic locking is undesirable.  *Patches
 try to changes to complex types, arrays, and maps.  The the merge is performed is defined on a per-field basis.*

### Read

Reads come in several flavors:

- Get
- List
- Watch

**Get**

Get will retrieve a specific resource object by name.

**List**

List will retrieve all resource objects and can be filtered by a selector query.

**Watch**

Watch will stream results for an object as they are updated.  Watch is useful for taking action in response to changes
to api objects.

### Delete

Delete will delete a resource.  Depending on the specific resource, child objects may or may not be deleted.  See
notes on specific resource objects for details.

**Other operations**

Each resource may define additional operations specific to it

### Rollback

Rollback a PodTemplate to a previous version.

### Read / Write Scale

Read or Update the number of replicas for the given resource.

### Read / Write Status

Read or Update the Status for a resource object.  The Status can only changed through these update operations.

### List for all namespaces

List resources across namespaces.  Normal list operations restrict results to the specified namespace.

### Watch List

------------
