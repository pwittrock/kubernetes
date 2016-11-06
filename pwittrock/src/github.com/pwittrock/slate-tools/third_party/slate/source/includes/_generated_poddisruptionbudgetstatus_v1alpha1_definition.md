## PodDisruptionBudgetStatus v1alpha1

Group        | Version     | Kind
------------ | ---------- | -----------
Policy | v1alpha1 | PodDisruptionBudgetStatus

> Example yaml coming soon...



PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.

<aside class="notice">
Appears In  <a href="#poddisruptionbudget-v1alpha1">PodDisruptionBudget</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
currentHealthy | integer | current number of healthy pods
desiredHealthy | integer | minimum desired number of healthy pods
disruptionAllowed | boolean | Whether or not a disruption is currently allowed.
expectedPods | integer | total number of pods counted by this disruption budget

