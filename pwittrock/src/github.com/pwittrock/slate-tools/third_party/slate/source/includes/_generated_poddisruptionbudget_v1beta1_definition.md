## PodDisruptionBudget v1beta1

Group        | Version     | Kind
------------ | ---------- | -----------
Policy | v1beta1 | PodDisruptionBudget

> Example yaml coming soon...



PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods

<aside class="notice">
Appears In  <a href="#poddisruptionbudgetlist-v1beta1">PodDisruptionBudgetList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | 
spec | [PodDisruptionBudgetSpec](#poddisruptionbudgetspec-v1beta1) | Specification of the desired behavior of the PodDisruptionBudget.
status | [PodDisruptionBudgetStatus](#poddisruptionbudgetstatus-v1beta1) | Most recently observed status of the PodDisruptionBudget.

