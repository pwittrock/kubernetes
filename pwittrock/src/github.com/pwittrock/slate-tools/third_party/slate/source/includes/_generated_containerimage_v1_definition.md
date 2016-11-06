## ContainerImage v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | ContainerImage

> Example yaml coming soon...



Describe a container image

<aside class="notice">
Appears In  <a href="#nodestatus-v1">NodeStatus</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
names | string array | Names by which this image is known. e.g. ["gcr.io/google_containers/hyperkube:v1.0.7", "dockerhub.io/google_containers/hyperkube:v1.0.7"]
sizeBytes | integer | The size of the image in bytes.

