## *CertificateSigningRequestStatus v1alpha1*

> Example yaml coming soon...





<aside class="notice">
Appears In  <a href="#certificatesigningrequest-v1alpha1">CertificateSigningRequest</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
conditions | [CertificateSigningRequestCondition](#certificatesigningrequestcondition-v1alpha1) array | Conditions applied to the request, such as approval or denial.
certificate | string | If request was approved, the controller will place the issued certificate here.

