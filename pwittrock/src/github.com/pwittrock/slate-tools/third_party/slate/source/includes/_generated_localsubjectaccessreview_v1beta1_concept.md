

-----------
# LocalSubjectAccessReview v1beta1

Group        | Version     | Kind
------------ | ---------- | -----------
Authorization | v1beta1 | LocalSubjectAccessReview







> Example yaml coming soon...


LocalSubjectAccessReview checks whether or not a user or group can perform an action in a given namespace. Having a namespace scoped resource makes it much easier to grant namespace scoped policy that includes permissions checking.



Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | 
spec | [SubjectAccessReviewSpec](#subjectaccessreviewspec-v1beta1) | Spec holds information about the request being evaluated.  spec.namespace must be equal to the namespace you made the request against.  If empty, it is defaulted.
status | [SubjectAccessReviewStatus](#subjectaccessreviewstatus-v1beta1) | Status is filled in by the server and indicates whether the request is allowed or not





## <strong>Write Operations</strong>

See supported operations below...

## Create

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



create a LocalSubjectAccessReview

### HTTP Request

`POST /apis/authorization.k8s.io/v1beta1/namespaces/{namespace}/localsubjectaccessreviews`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [LocalSubjectAccessReview](#localsubjectaccessreview-v1beta1) | 
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [LocalSubjectAccessReview](#localsubjectaccessreview-v1beta1) | OK




