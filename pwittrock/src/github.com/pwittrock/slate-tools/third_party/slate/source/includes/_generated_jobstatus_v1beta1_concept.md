

-----------
# JobStatus v1beta1



Group        | Version     | Kind
------------ | ---------- | -----------
Extensions | v1beta1 | JobStatus




<aside class="notice">Other api versions of this object exist: <a href="#jobstatus-v1">v1</a> </aside>


JobStatus represents the current state of a Job.

<aside class="notice">
Appears In <a href="#job-v1beta1">Job</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
active | integer | Active is the number of actively running pods.
completionTime | [Time](#time-unversioned) | CompletionTime represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
conditions | [JobCondition](#jobcondition-v1beta1) array | Conditions represent the latest available observations of an object's current state. More info: http://kubernetes.io/docs/user-guide/jobs
failed | integer | Failed is the number of pods which reached Phase Failed.
startTime | [Time](#time-unversioned) | StartTime represents time when the job was acknowledged by the Job Manager. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
succeeded | integer | Succeeded is the number of pods which reached Phase Succeeded.





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



replace status of the specified Job

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/jobs/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Job
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Job](#job-v1beta1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Job](#job-v1beta1) | OK


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



partially update status of the specified Job

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/jobs/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Job
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Patch](#patch-unversioned) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Job](#job-v1beta1) | OK



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



read status of the specified Job

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/jobs/{name}/status`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Job
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Job](#job-v1beta1) | OK




