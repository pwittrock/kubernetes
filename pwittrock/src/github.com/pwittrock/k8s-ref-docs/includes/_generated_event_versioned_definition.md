## *Event versioned*

> Example yaml coming soon...



Event represents a single event to a watched resource.



Field        | Schema     | Description
------------ | ---------- | -----------
object | [RawExtension](#rawextension-runtime) | Object is:
 * If Type is Added or Modified: the new state of the object.
 * If Type is Deleted: the state of the object immediately before deletion.
 * If Type is Error: *api.Status is recommended; other types may make sense
   depending on context.
type | string | 

