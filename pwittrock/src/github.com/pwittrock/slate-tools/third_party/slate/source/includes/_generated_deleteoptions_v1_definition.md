## DeleteOptions v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | DeleteOptions

> Example yaml coming soon...



DeleteOptions may be provided when deleting an API object

<aside class="notice">
Appears In  <a href="#eviction-v1beta1">Eviction</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
gracePeriodSeconds | integer | The duration in seconds before the object should be deleted. Value must be non-negative integer. The value zero indicates delete immediately. If this value is nil, the default grace period for the specified type will be used. Defaults to a per object value if not specified. zero means delete immediately.
orphanDependents | boolean | Should the dependent objects be orphaned. If true/false, the "orphan" finalizer will be added to/removed from the object's finalizers list.
preconditions | [Preconditions](#preconditions-v1) | Must be fulfilled before a deletion is carried out. If not possible, a 409 Conflict status will be returned.

