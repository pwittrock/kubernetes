## *ClusterRoleBinding v1alpha1*

> Example yaml coming soon...



ClusterRoleBinding references a ClusterRole, but not contain it.  It can reference a ClusterRole in the global namespace, and adds who information via Subject.

<aside class="notice">
Appears In  <a href="#clusterrolebindinglist-v1alpha1">ClusterRoleBindingList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
roleRef | [RoleRef](#roleref-v1alpha1) | RoleRef can only reference a ClusterRole in the global namespace. If the RoleRef cannot be resolved, the Authorizer must return an error.
subjects | [Subject](#subject-v1alpha1) array | Subjects holds references to the objects the role applies to.
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata.

