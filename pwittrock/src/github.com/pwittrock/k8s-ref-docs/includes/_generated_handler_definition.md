## *Handler v1*

> Example yaml coming soon...



Handler defines a specific action that should be taken

<aside class="notice">
Appears In  <a href="#lifecycle-v1">Lifecycle</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
exec | [ExecAction](#execaction-v1) | One and only one of the following should be specified. Exec specifies the action to take.
httpGet | [HTTPGetAction](#httpgetaction-v1) | HTTPGet specifies the http request to perform.
tcpSocket | [TCPSocketAction](#tcpsocketaction-v1) | TCPSocket specifies an action involving a TCP port. TCP hooks not yet supported

