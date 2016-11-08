## *ThirdPartyResource v1beta1*

> Example yaml coming soon...



A ThirdPartyResource is a generic representation of a resource, it is used by add-ons and plugins to add new resource types to the API.  It consists of one or more Versions of the api.

<aside class="notice">
Appears In  <a href="#thirdpartyresourcelist-v1beta1">ThirdPartyResourceList</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
versions | [APIVersion](#apiversion-v1beta1) array | Versions are versions for this third party object
description | string | Description is the description of this object.
metadata | [ObjectMeta](#objectmeta-v1) | Standard object metadata

