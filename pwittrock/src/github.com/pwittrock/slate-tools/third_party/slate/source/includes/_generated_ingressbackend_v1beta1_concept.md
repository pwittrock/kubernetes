

-----------
# IngressBackend v1beta1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1beta1 | IngressBackend







> Example yaml coming soon...


IngressBackend describes all endpoints for a given service and port.

<aside class="notice">
Appears In <a href="#ingressspec-v1beta1">IngressSpec</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
serviceName | string | Specifies the name of the referenced service.
servicePort | [IntOrString](#intorstring-intstr) | Specifies the port of the referenced service.






