

-----------
# LabelSelector v1



Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | LabelSelector




<aside class="notice">Other api versions of this object exist: <a href="#labelselector-v1beta1">v1beta1</a> <a href="#labelselector-unversioned">unversioned</a> </aside>


A label selector is a label query over a set of resources. The result of matchLabels and matchExpressions are ANDed. An empty label selector matches all objects. A null label selector matches no objects.

<aside class="notice">
Appears In <a href="#jobspec-v1">JobSpec</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
matchExpressions | [LabelSelectorRequirement](#labelselectorrequirement-v1) array | matchExpressions is a list of label selector requirements. The requirements are ANDed.
matchLabels | object | matchLabels is a map of {key,value} pairs. A single {key,value} in the matchLabels map is equivalent to an element of matchExpressions, whose key field is "key", the operator is "In", and the values array contains only "value". The requirements are ANDed.






