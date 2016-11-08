

-----------

# Job v1




<aside class="notice">Older api versions of this object exist: <a href="#job-v1beta1">v1beta1</a> </aside>

> Example yaml coming soon...


Job represents the configuration of a single job.

<aside class="notice">
Appears In <a href="#joblist-v1">JobList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [JobSpec](#jobspec-v1) | Spec is a structure defining the expected behavior of a job. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status
status | [JobStatus](#jobstatus-v1) | Status is a structure describing current status of a job. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status


### JobSpec v1

<aside class="notice">
Appears In <a href="#job-v1">Job</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
activeDeadlineSeconds | integer | Optional duration in seconds relative to the startTime that the job may be active before the system tries to terminate it; value must be positive integer
completions | integer | Completions specifies the desired number of successfully finished pods the job should be run with.  Setting to nil means that the success of any pod signals the success of all pods, and allows parallelism to have any positive value.  Setting to 1 means that parallelism is limited to 1 and the success of that pod signals the success of the job. More info: http://kubernetes.io/docs/user-guide/jobs
manualSelector | boolean | ManualSelector controls generation of pod labels and pod selectors. Leave `manualSelector` unset unless you are certain what you are doing. When false or unset, the system pick labels unique to this job and appends those labels to the pod template.  When true, the user is responsible for picking unique labels and specifying the selector.  Failure to pick a unique label may cause this and other jobs to not function correctly.  However, You may see `manualSelector=true` in jobs that were created with the old `extensions/v1beta1` API. More info: http://releases.k8s.io/HEAD/docs/design/selector-generation.md
parallelism | integer | Parallelism specifies the maximum desired number of pods the job should run at any given time. The actual number of pods running in steady state will be less than this number when ((.spec.completions - .status.successful) < .spec.parallelism), i.e. when the work left to do is less than max parallelism. More info: http://kubernetes.io/docs/user-guide/jobs
selector | [LabelSelector](#labelselector-v1) | Selector is a label query over pods that should match the pod count. Normally, the system sets this field for you. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors
template | [PodTemplateSpec](#podtemplatespec-v1) | Template is the object that describes the pod that will be created when executing a job. More info: http://kubernetes.io/docs/user-guide/jobs

### JobStatus v1

<aside class="notice">
Appears In <a href="#job-v1">Job</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
active | integer | Active is the number of actively running pods.
completionTime | [Time](#time-unversioned) | CompletionTime represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
conditions | [JobCondition](#jobcondition-v1) array | Conditions represent the latest available observations of an object's current state. More info: http://kubernetes.io/docs/user-guide/jobs
failed | integer | Failed is the number of pods which reached Phase Failed.
startTime | [Time](#time-unversioned) | StartTime represents time when the job was acknowledged by the Job Manager. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
succeeded | integer | Succeeded is the number of pods which reached Phase Succeeded.

### JobList v1



Field        | Schema     | Description
------------ | ---------- | -----------
items | [Job](#job-v1) array | Items is the list of Job.
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata




## <strong>Write Operations</strong>

See supported operations below...

## Create

> Examples using curl coming soon...

create a Job

### HTTP Request

`POST /apis/extensions/v1beta1/namespaces/{namespace}/jobs`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Job](#job-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Job](#job-v1) | OK


## Replace

> Examples using curl coming soon...

replace the specified Job

### HTTP Request

`PUT /apis/extensions/v1beta1/namespaces/{namespace}/jobs/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Job
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
body | [Job](#job-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Job](#job-v1) | OK


## Patch

> Examples using curl coming soon...

partially update the specified Job

### HTTP Request

`PATCH /apis/extensions/v1beta1/namespaces/{namespace}/jobs/{name}`

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
200 | [Job](#job-v1) | OK


## Delete

> Examples using curl coming soon...

delete a Job

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/jobs/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Job
namespace |  | object name and auth scope, such as for teams and projects
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

> Examples using curl coming soon...

read the specified Job

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/jobs/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
name |  | name of the Job
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.

### Query Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
exact |  | Should the export be exact.  Exact export maintains cluster-specific fields like 'Namespace'
export |  | Should this value be exported.  Export strips fields that a user can not specify.

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Job](#job-v1) | OK


## List

> Examples using curl coming soon...

list or watch objects of kind Job

### HTTP Request

`GET /apis/extensions/v1beta1/namespaces/{namespace}/jobs`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
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
200 | [JobList](#joblist-v1) | OK


## Watch

> Examples using curl coming soon...

watch changes to an object of kind Job

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/jobs/{name}`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
name |  | name of the Job
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-v1) | OK



## <strong>Status & Collection Operations</strong>

See supported operations below...

## Patch Status

> Examples using curl coming soon...

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
200 | [Job](#job-v1) | OK


## Read Status

> Examples using curl coming soon...

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
200 | [Job](#job-v1) | OK


## Replace Status

> Examples using curl coming soon...

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
body | [Job](#job-v1) | 

### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Job](#job-v1) | OK


## Delete Collection

> Examples using curl coming soon...

delete collection of Job

### HTTP Request

`DELETE /apis/extensions/v1beta1/namespaces/{namespace}/jobs`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
namespace |  | object name and auth scope, such as for teams and projects
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


## List All Namespaces

> Examples using curl coming soon...

list or watch objects of kind Job

### HTTP Request

`GET /apis/extensions/v1beta1/jobs`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [JobList](#joblist-v1) | OK


## Watch List

> Examples using curl coming soon...

watch individual changes to a list of Job

### HTTP Request

`GET /apis/extensions/v1beta1/watch/namespaces/{namespace}/jobs`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
namespace |  | object name and auth scope, such as for teams and projects
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-v1) | OK


## Watch List All Namespaces

> Examples using curl coming soon...

watch individual changes to a list of Job

### HTTP Request

`GET /apis/extensions/v1beta1/watch/jobs`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-v1) | OK




