

-----------
# OwnerReference v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | OwnerReference







OwnerReference contains enough information to let you identify an owning object. Currently, an owning object must be in the same namespace, so there is no namespace field.

<aside class="notice">
Appears In <a href="#objectmeta-v1">ObjectMeta</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
apiVersion | string | API version of the referent.
controller | boolean | If true, this reference points to the managing controller.
kind | string | Kind of the referent. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#types-kinds
name | string | Name of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#names
uid | string | UID of the referent. More info: http://kubernetes.io/docs/user-guide/identifiers#uids






