

-----------
# Probe v1

Group        | Version     | Kind
------------ | ---------- | -----------
Core | v1 | Probe







> Example yaml coming soon...


Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.

<aside class="notice">
Appears In <a href="#container-v1">Container</a> </aside>

Field        | Schema     | Description
------------ | ---------- | -----------
failureThreshold | integer | Minimum consecutive failures for the probe to be considered failed after having succeeded. Defaults to 3. Minimum value is 1.
initialDelaySeconds | integer | Number of seconds after the container has started before liveness probes are initiated. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes
periodSeconds | integer | How often (in seconds) to perform the probe. Default to 10 seconds. Minimum value is 1.
successThreshold | integer | Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1. Must be 1 for liveness. Minimum value is 1.
timeoutSeconds | integer | Number of seconds after which the probe times out. Defaults to 1 second. Minimum value is 1. More info: http://kubernetes.io/docs/user-guide/pod-states#container-probes






