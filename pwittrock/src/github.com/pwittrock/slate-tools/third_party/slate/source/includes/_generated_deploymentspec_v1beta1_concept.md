

-----------
# DeploymentSpec v1beta1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1beta1 | DeploymentSpec







DeploymentSpec is the specification of the desired behavior of the Deployment.

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






