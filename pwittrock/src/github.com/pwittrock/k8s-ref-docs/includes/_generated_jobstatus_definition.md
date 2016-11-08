## *JobStatus v1*

> Example yaml coming soon...

<aside class="notice">Older api versions of this object exist: <a href="#jobstatus-v1beta1">v1beta1</a> </aside>

JobStatus represents the current state of a Job.

<aside class="notice">
Appears In  <a href="#job-v1">Job</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
active | integer | Active is the number of actively running pods.
completionTime | [Time](#time-unversioned) | CompletionTime represents time when the job was completed. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
conditions | [JobCondition](#jobcondition-v1) array | Conditions represent the latest available observations of an object's current state. More info: http://kubernetes.io/docs/user-guide/jobs
failed | integer | Failed is the number of pods which reached Phase Failed.
startTime | [Time](#time-unversioned) | StartTime represents time when the job was acknowledged by the Job Manager. It is not guaranteed to be set in happens-before order across separate operations. It is represented in RFC3339 form and is in UTC.
succeeded | integer | Succeeded is the number of pods which reached Phase Succeeded.

