

-----------
# TokenReview v1beta1



Group        | Version     | Kind
------------ | ---------- | -----------
Authentication | v1beta1 | TokenReview







TokenReview attempts to authenticate a token to a known user. Note: TokenReview requests may be cached by the webhook token authenticator plugin in the kube-apiserver.



Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | 
spec | [TokenReviewSpec](#tokenreviewspec-v1beta1) | Spec holds information about the request being evaluated
status | [TokenReviewStatus](#tokenreviewstatus-v1beta1) | Status is filled in by the server and indicates whether the request can be authenticated.


### TokenReviewSpec v1beta1

<aside class="notice">
Appears In <a href="#tokenreview-v1beta1">TokenReview</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
token | string | Token is the opaque bearer token.

### TokenReviewStatus v1beta1

<aside class="notice">
Appears In <a href="#tokenreview-v1beta1">TokenReview</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
authenticated | boolean | Authenticated indicates that the token was associated with a known user.
error | string | Error indicates that the token couldn't be checked
user | [UserInfo](#userinfo-v1beta1) | User is the UserInfo associated with the provided token.




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



create a TokenReview

### HTTP Request

`POST /apis/authentication.k8s.io/v1beta1/tokenreviews`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [TokenReview](#tokenreview-v1beta1) | 
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [TokenReview](#tokenreview-v1beta1) | OK




