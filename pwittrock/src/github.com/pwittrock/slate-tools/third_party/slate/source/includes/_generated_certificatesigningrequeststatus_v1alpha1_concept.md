

-----------
# CertificateSigningRequestStatus v1alpha1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1alpha1 | CertificateSigningRequestStatus







> Example yaml coming soon...




<aside class="notice">
Appears In <a href="#certificatesigningrequest-v1alpha1">CertificateSigningRequest</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
certificate | string | If request was approved, the controller will place the issued certificate here.
conditions | [CertificateSigningRequestCondition](#certificatesigningrequestcondition-v1alpha1) array | Conditions applied to the request, such as approval or denial.






