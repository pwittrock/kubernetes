## *TokenReview v1beta1*

> Example yaml coming soon...



TokenReview attempts to authenticate a token to a known user. Note: TokenReview requests may be cached by the webhook token authenticator plugin in the kube-apiserver.



Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | 
spec | [TokenReviewSpec](#tokenreviewspec-v1beta1) | Spec holds information about the request being evaluated
status | [TokenReviewStatus](#tokenreviewstatus-v1beta1) | Status is filled in by the server and indicates whether the request can be authenticated.

