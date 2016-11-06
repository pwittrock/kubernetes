## EndpointPort v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | EndpointPort

> Example yaml coming soon...



EndpointPort is a tuple that describes a single port.

<aside class="notice">
Appears In  <a href="#endpointsubset-v1">EndpointSubset</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
name | string | The name of this port (corresponds to ServicePort.Name). Must be a DNS_LABEL. Optional only if one port is defined.
port | integer | The port number of the endpoint.
protocol | string | The IP protocol for this port. Must be UDP or TCP. Default is TCP.

