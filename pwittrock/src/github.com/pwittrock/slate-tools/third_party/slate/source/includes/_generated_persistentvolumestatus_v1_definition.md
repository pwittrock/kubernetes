## PersistentVolumeStatus v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | PersistentVolumeStatus

> Example yaml coming soon...



PersistentVolumeStatus is the current status of a persistent volume.

<aside class="notice">
Appears In  <a href="#persistentvolume-v1">PersistentVolume</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
message | string | A human-readable message indicating details about why the volume is in this state.
phase | string | Phase indicates if a volume is available, bound to a claim, or released by a claim. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#phase
reason | string | Reason is a brief CamelCase string that describes any failure and is meant for machine parsing and tidy display in the CLI.

