## *Event v1*

> Example yaml coming soon...

<aside class="notice">Older api versions of this object exist: <a href="#event-versioned">versioned</a> </aside>

Event is a report of an event somewhere in the cluster.

<aside class="notice">
Appears In  <a href="#eventlist-v1">EventList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
count | integer | The number of times this event has occurred.
involvedObject | [ObjectReference](#objectreference-v1) | The object that this event is about.
lastTimestamp | [Time](#time-unversioned) | The time at which the most recent occurrence of this event was recorded.
metadata | [ObjectMeta](#objectmeta-v1) | Standard object's metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
reason | string | This should be a short, machine understandable string that gives the reason for the transition into the object's current status.
firstTimestamp | [Time](#time-unversioned) | The time at which the event was first recorded. (Time of server receipt is in TypeMeta.)
message | string | A human-readable description of the status of this operation.
source | [EventSource](#eventsource-v1) | The component reporting this event. Should be a short machine understandable string.
type | string | Type of this event (Normal, Warning), new types could be added in the future

