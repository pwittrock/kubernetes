

-----------
# ResourceAttributes v1beta1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1beta1 | ResourceAttributes







ResourceAttributes includes the authorization attributes available for resource requests to the Authorizer interface

<aside class="notice">
Appears In <a href="#selfsubjectaccessreviewspec-v1beta1">SelfSubjectAccessReviewSpec</a> <a href="#subjectaccessreviewspec-v1beta1">SubjectAccessReviewSpec</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
group | string | Group is the API Group of the Resource.  "*" means all.
name | string | Name is the name of the resource being requested for a "get" or deleted for a "delete". "" (empty) means all.
namespace | string | Namespace is the namespace of the action being requested.  Currently, there is no distinction between no namespace and all namespaces "" (empty) is defaulted for LocalSubjectAccessReviews "" (empty) is empty for cluster-scoped resources "" (empty) means "all" for namespace scoped resources from a SubjectAccessReview or SelfSubjectAccessReview
resource | string | Resource is one of the existing resource types.  "*" means all.
subresource | string | Subresource is one of the existing resource types.  "" means none.
verb | string | Verb is a kubernetes resource API verb, like: get, list, watch, create, update, delete, proxy.  "*" means all.
version | string | Version is the API Version of the Resource.  "*" means all.






