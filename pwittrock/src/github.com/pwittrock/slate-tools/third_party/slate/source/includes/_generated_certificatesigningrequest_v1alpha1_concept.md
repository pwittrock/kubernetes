

-----------
# CertificateSigningRequest v1alpha1



Group        | Version     | Kind
------------ | ---------- | -----------
Certificates | v1alpha1 | CertificateSigningRequest







Describes a certificate signing request

<aside class="notice">
Appears In <a href="#certificatesigningrequestlist-v1alpha1">CertificateSigningRequestList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | 
spec | [CertificateSigningRequestSpec](#certificatesigningrequestspec-v1alpha1) | The certificate request itself and any additional information.
status | [CertificateSigningRequestStatus](#certificatesigningrequeststatus-v1alpha1) | Derived information about the request.


### CertificateSigningRequestSpec v1alpha1

<aside class="notice">
Appears In <a href="#certificatesigningrequest-v1alpha1">CertificateSigningRequest</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
groups | string array | 
request | string | Base64-encoded PKCS#10 CSR data
uid | string | 
username | string | Information about the requesting user (if relevant) See user.Info interface for details

### CertificateSigningRequestStatus v1alpha1

<aside class="notice">
Appears In <a href="#certificatesigningrequest-v1alpha1">CertificateSigningRequest</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
certificate | string | If request was approved, the controller will place the issued certificate here.
conditions | [CertificateSigningRequestCondition](#certificatesigningrequestcondition-v1alpha1) array | Conditions applied to the request, such as approval or denial.

### CertificateSigningRequestList v1alpha1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [CertificateSigningRequest](#certificatesigningrequest-v1alpha1) array | 
metadata | [ListMeta](#listmeta-unversioned) | 




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



create a CertificateSigningRequest

### HTTP Request

`POST /apis/certificates.k8s.io/v1alpha1/certificatesigningrequests`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [CertificateSigningRequest](#certificatesigningrequest-v1alpha1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [CertificateSigningRequest](#certificatesigningrequest-v1alpha1) | OK


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



replace the specified CertificateSigningRequest

### HTTP Request

`PUT /apis/certificates.k8s.io/v1alpha1/certificatesigningrequests/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the CertificateSigningRequest
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [CertificateSigningRequest](#certificatesigningrequest-v1alpha1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [CertificateSigningRequest](#certificatesigningrequest-v1alpha1) | OK


## Patch

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



partially update the specified CertificateSigningRequest

### HTTP Request

`PATCH /apis/certificates.k8s.io/v1alpha1/certificatesigningrequests/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the CertificateSigningRequest
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [CertificateSigningRequest](#certificatesigningrequest-v1alpha1) | OK


## Delete

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



delete a CertificateSigningRequest

### HTTP Request

`DELETE /apis/certificates.k8s.io/v1alpha1/certificatesigningrequests/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the CertificateSigningRequest
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [DeleteOptions](#deleteoptions-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Status](#status-unversioned) | OK



## <strong>Read Operations</strong>

See supported operations below...

## Read

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



read the specified CertificateSigningRequest

### HTTP Request

`GET /apis/certificates.k8s.io/v1alpha1/certificatesigningrequests/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the CertificateSigningRequest
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
exact |  | Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'
export |  | Should this value be exported.  Export strips fields that a user can not specify.

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [CertificateSigningRequest](#certificatesigningrequest-v1alpha1) | OK


## List

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



list or watch objects of kind CertificateSigningRequest

### HTTP Request

`GET /apis/certificates.k8s.io/v1alpha1/certificatesigningrequests`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [CertificateSigningRequestList](#certificatesigningrequestlist-v1alpha1) | OK


## Watch

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



watch changes to an object of kind CertificateSigningRequest

### HTTP Request

`GET /apis/certificates.k8s.io/v1alpha1/watch/certificatesigningrequests/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the CertificateSigningRequest
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-versioned) | OK



## <strong>Status & Collection Operations</strong>

See supported operations below...

## Replace Status

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


## Delete Collection

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



delete collection of CertificateSigningRequest

### HTTP Request

`DELETE /apis/certificates.k8s.io/v1alpha1/certificatesigningrequests`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Status](#status-unversioned) | OK




