## *PersistentVolumeClaimSpec v1*

> Example yaml coming soon...



PersistentVolumeClaimSpec describes the common attributes of storage devices and allows a Source for provider-specific attributes

<aside class="notice">
Appears In  <a href="#persistentvolumeclaim-v1">PersistentVolumeClaim</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
accessModes | string array | AccessModes contains the desired access modes the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#access-modes-1
resources | [ResourceRequirements](#resourcerequirements-v1) | Resources represents the minimum resources the volume should have. More info: http://kubernetes.io/docs/user-guide/persistent-volumes#resources
selector | [LabelSelector](#labelselector-v1) | A label query over volumes to consider for binding.
volumeName | string | VolumeName is the binding reference to the PersistentVolume backing this claim.

