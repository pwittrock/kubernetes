

-----------
# CertificateSigningRequestStatus v1alpha1



Group        | Version     | Kind
------------ | ---------- | -----------
Certificates | v1alpha1 | CertificateSigningRequestStatus









<aside class="notice">
Appears In <a href="#certificatesigningrequest-v1alpha1">CertificateSigningRequest</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
certificate | string | If request was approved, the controller will place the issued certificate here.
conditions | [CertificateSigningRequestCondition](#certificatesigningrequestcondition-v1alpha1) array | Conditions applied to the request, such as approval or denial.





## <strong>Write Operations</strong>

See supported operations below...

## Replace

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



replace status of the specified CertificateSigningRequest

### HTTP Request

`PUT /apis/certificates.k8s.io/v1alpha1/certificatesigningrequests/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [CertificateSigningRequest](#certificatesigningrequest-v1alpha1) | 
name |  | name of the CertificateSigningRequest
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [CertificateSigningRequest](#certificatesigningrequest-v1alpha1) | OK




