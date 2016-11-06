## PetSet v1alpha1

Group        | Version     | Kind
------------ | ---------- | -----------
Apps | v1alpha1 | PetSet

> Example yaml coming soon...



PetSet represents a set of pods with consistent identities. Identities are defined as:
 - Network: A single stable DNS and hostname.
 - Storage: As many VolumeClaims as requested.
The PetSet guarantees that a given network identity will always map to the same storage identity. PetSet is currently in alpha and subject to change without notice.

<aside class="notice">
Appears In  <a href="#petsetlist-v1alpha1">PetSetList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | 
spec | [PetSetSpec](#petsetspec-v1alpha1) | Spec defines the desired identities of pets in this set.
status | [PetSetStatus](#petsetstatus-v1alpha1) | Status is the current status of Pets in this PetSet. This data may be out of date by some window of time.

