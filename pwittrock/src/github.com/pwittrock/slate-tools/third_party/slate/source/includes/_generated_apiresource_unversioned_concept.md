

-----------
# APIResource unversioned

Group        | Version     | Kind
------------ | ---------- | -----------
Core | unversioned | APIResource







> Example yaml coming soon...


APIResource specifies the name of a resource and whether it is namespaced.

<aside class="notice">
Appears In <a href="#apiresourcelist-unversioned">APIResourceList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
kind | string | kind is the kind for the resource (e.g. 'Foo' is the kind for a resource 'foo')
name | string | name is the name of the resource.
namespaced | boolean | namespaced indicates if a resource is namespaced or not.


### APIResourceList unversioned



Field        | Schema     | Description
------------ | ---------- | -----------
groupVersion | string | groupVersion is the group and version this APIResourceList is for.
resources | [APIResource](#apiresource-unversioned) array | resources contains the name of the resources and if they are namespaced.





