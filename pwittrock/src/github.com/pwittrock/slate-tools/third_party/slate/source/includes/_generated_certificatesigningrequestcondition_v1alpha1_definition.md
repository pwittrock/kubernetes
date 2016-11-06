## CertificateSigningRequestCondition v1alpha1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1alpha1 | CertificateSigningRequestCondition

> Example yaml coming soon...





<aside class="notice">
Appears In  <a href="#certificatesigningrequeststatus-v1alpha1">CertificateSigningRequestStatus</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
lastUpdateTime | [Time](#time-unversioned) | timestamp for the last update to this condition
message | string | human readable message with details about the request state
reason | string | brief reason for the request state
type | string | request approval state, currently Approved or Denied.

