

-----------
# ThirdPartyResourceList v1beta1

Group        | Version     | Kind
------------ | ---------- | -----------
Extensions | v1beta1 | ThirdPartyResourceList







> Example yaml coming soon...


ThirdPartyResourceList is a list of ThirdPartyResources.



Field        | Schema     | Description
------------ | ---------- | -----------
items | [ThirdPartyResource](#thirdpartyresource-v1beta1) array | Items is the list of ThirdPartyResources.
metadata | [ListMeta](#listmeta-unversioned) | Standard list metadata.





## <strong>Read Operations</strong>

See supported operations below...

## Watch

> Execute

```shell



```



```yaml



```

> Returns

```shell



```


```yaml



```



watch individual changes to a list of ThirdPartyResource

### HTTP Request

`GET /apis/extensions/v1beta1/watch/thirdpartyresources`

### Path Parameters

Parameter    | Schema     | Description
------------ | ---------- | -----------
fieldSelector |  | A selector to restrict the list of returned objects by their fields. Defaults to everything.
labelSelector |  | A selector to restrict the list of returned objects by their labels. Defaults to everything.
pretty |  | If 'true', then the output is pretty printed.
resourceVersion |  | When specified with a watch call, shows changes that occur after that particular version of a resource. Defaults to changes from the beginning of history.
timeoutSeconds |  | Timeout for the list/watch call.
watch |  | Watch for changes to the described resources and return them as a stream of add, update, and remove notifications. Specify resourceVersion.


### Response

Code         | Schema     | Description
------------ | ---------- | -----------
200 | [Event](#event-versioned) | OK




