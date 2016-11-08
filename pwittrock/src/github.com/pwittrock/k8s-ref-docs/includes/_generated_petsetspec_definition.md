## *PetSetSpec v1alpha1*

> Example yaml coming soon...



A PetSetSpec is the specification of a PetSet.

<aside class="notice">
Appears In  <a href="#petset-v1alpha1">PetSet</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
serviceName | string | ServiceName is the name of the service that governs this PetSet. This service must exist before the PetSet, and is responsible for the network identity of the set. Pets get DNS/hostnames that follow the pattern: pet-specific-string.serviceName.default.svc.cluster.local where "pet-specific-string" is managed by the PetSet controller.
template | [PodTemplateSpec](#podtemplatespec-v1) | Template is the object that describes the pod that will be created if insufficient replicas are detected. Each pod stamped out by the PetSet will fulfill this Template, but have a unique identity from the rest of the PetSet.
volumeClaimTemplates | [PersistentVolumeClaim](#persistentvolumeclaim-v1) array | VolumeClaimTemplates is a list of claims that pets are allowed to reference. The PetSet controller is responsible for mapping network identities to claims in a way that maintains the identity of a pet. Every claim in this list must have at least one matching (by name) volumeMount in one container in the template. A claim in this list takes precedence over any volumes in the template, with the same name.
replicas | integer | Replicas is the desired number of replicas of the given Template. These are replicas in the sense that they are instantiations of the same Template, but individual replicas also have a consistent identity. If unspecified, defaults to 1.
selector | [LabelSelector](#labelselector-v1) | Selector is a label query over pods that should match the replica count. If empty, defaulted to labels on the pod template. More info: http://kubernetes.io/docs/user-guide/labels#label-selectors

