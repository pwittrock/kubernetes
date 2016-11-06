

-----------
# ContainerStateTerminated v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ContainerStateTerminated







> Example yaml coming soon...


ContainerStateTerminated is a terminated state of a container.

<aside class="notice">
Appears In <a href="#containerstate-v1">ContainerState</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
containerID | string | Container's ID in the format 'docker://<container_id>'
exitCode | integer | Exit status from the last termination of the container
finishedAt | [Time](#time-unversioned) | Time at which the container last terminated
message | string | Message regarding the last termination of the container
reason | string | (brief) reason from the last termination of the container
signal | integer | Signal from the last termination of the container
startedAt | [Time](#time-unversioned) | Time at which previous execution of the container started






