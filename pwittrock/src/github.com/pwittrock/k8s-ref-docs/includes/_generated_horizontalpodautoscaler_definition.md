## *HorizontalPodAutoscaler v1*

> Example yaml coming soon...

<aside class="notice">Older api versions of this object exist: <a href="#horizontalpodautoscaler-v1beta1">v1beta1</a> </aside>

configuration of a horizontal pod autoscaler.

<aside class="notice">
Appears In  <a href="#horizontalpodautoscalerlist-v1">HorizontalPodAutoscalerList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
metadata | [ObjectMeta](#objectmeta-v1) | Standard object metadata. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#metadata
spec | [HorizontalPodAutoscalerSpec](#horizontalpodautoscalerspec-v1) | behaviour of autoscaler. More info: http://releases.k8s.io/HEAD/docs/devel/api-conventions.md#spec-and-status.
status | [HorizontalPodAutoscalerStatus](#horizontalpodautoscalerstatus-v1) | current information about the autoscaler.

