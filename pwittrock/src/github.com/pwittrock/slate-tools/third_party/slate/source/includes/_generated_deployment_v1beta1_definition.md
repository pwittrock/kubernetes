## Deployment v1beta1

Group        | Version     | Kind
------------ | ---------- | -----------
Extensions | v1beta1 | Deployment

> Example yaml coming soon...



Deployment enables declarative updates for Pods and ReplicaSets.

<aside class="notice">
Appears In  <a href="#deploymentlist-v1beta1">DeploymentList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object metadata.
spec | [DeploymentSpec](#deploymentspec-v1beta1) | Specification of the desired behavior of the Deployment.
status | [DeploymentStatus](#deploymentstatus-v1beta1) | Most recently observed status of the Deployment.

