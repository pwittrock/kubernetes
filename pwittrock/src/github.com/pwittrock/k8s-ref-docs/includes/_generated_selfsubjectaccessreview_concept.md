

-----------

# SelfSubjectAccessReview v1beta1






> Example yaml coming soon...


SelfSubjectAccessReview checks whether or the current user can perform an action.  Not filling in a spec.namespace means "in all namespaces".  Self is a special case, because users should always be able to check whether they can perform an action



Field        | Schema     | Description
------------ | ---------- | -----------
spec | [SelfSubjectAccessReviewSpec](#selfsubjectaccessreviewspec-v1beta1) | Spec holds information about the request being evaluated.  user and groups must be empty
status | [SubjectAccessReviewStatus](#subjectaccessreviewstatus-v1beta1) | Status is filled in by the server and indicates whether the request is allowed or not
metadata | [ObjectMeta](#objectmeta-v1) | 


### SelfSubjectAccessReviewSpec v1beta1

<aside class="notice">
Appears In <a href="#selfsubjectaccessreview-v1beta1">SelfSubjectAccessReview</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
nonResourceAttributes | [NonResourceAttributes](#nonresourceattributes-v1beta1) | NonResourceAttributes describes information for a non-resource access request
resourceAttributes | [ResourceAttributes](#resourceattributes-v1beta1) | ResourceAuthorizationAttributes describes information for a resource access request




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Examples using curl coming soon...

create a SelfSubjectAccessReview

### HTTP Request

`POST /apis/authorization.k8s.io/v1beta1/selfsubjectaccessreviews`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [SelfSubjectAccessReview](#selfsubjectaccessreview-v1beta1) | 
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [SelfSubjectAccessReview](#selfsubjectaccessreview-v1beta1) | OK




