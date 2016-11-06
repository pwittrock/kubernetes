

-----------
# Eviction v1alpha1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1alpha1 | Eviction







> Example yaml coming soon...


Eviction evicts a pod from its node subject to certain policies and safety constraints. This is a subresource of Pod.  A request to cause such an eviction is created by POSTing to .../pods/<pod name>/evictions.



Field        | Schema     | Description
------------ | ---------- | -----------
deleteOptions | [DeleteOptions](#deleteoptions-v1) | DeleteOptions may be provided
metadata | [ObjectMeta](#objectmeta-v1) | ObjectMeta describes the pod that is being evicted.






