# Getting Started on K8S

## Download docker
> TODO: Add instructions on how to do this

## Start a local cluster for development

Open a shell and enter

```
$ docker run --net host -t -i -v /var/run/docker.sock:/var/run/docker.sock -v $(which docker):/bin/docker pwittrock/k8sguide
```

This will bootstrap the key components of a k8s cluster locally, and provide a shell for interacting with the cluster.

TODO: currently this does not bootstrap dns or other required cluster add ons.  We need to make the cluster fully functional.

### Check that the cluster is healthy

```
$ kubectl get nodes
NAME        LABELS                             STATUS    AGE
127.0.0.1   kubernetes.io/hostname=127.0.0.1   Ready     14s
```


This will print out a list of nodes (machines) available for scheduling work.  Our cluster has only 1 node available
for scheduling work.  We will set up a cluster with multiple nodes later, but for now lets look at what we can do
with our single machine.

## Work with Pods

### Schedule work in a *Pod*

```
$ kubectl run --restart=Never nginx --image=nginx 
```

This will schedule a single instance of work called a *Pod* to be run on a node.
Since we only have 1 node in the cluster, there is only place it will be
scheduled.  In a multi-node cluster, the pod could be scheduled on any node with available resources.

Breaking down the command

```
run
```

Tells k8s to schedule work to be run

```
--restart=Never
```


Tells k8s not to monitor pod to be restarted if it fails.  Monitoring pods is done through a *Replication Controller*
which will be talked about in detail later, but for now you can ignore it.

```
nginx
```

This is the name k8s will give the pod

```
--image=nginx
```

This is the docker image to start


### Check the state of the pod

```
$ kubectl get pods
NAME      READY     STATUS              RESTARTS   AGE
nginx     0/1       ContainerCreating   0          9s
```


This shows us a list of all of the pods running in our namespace.  The *Pod* is in the process of creating the container.


```
$ kubectl get pods
NAME      READY     STATUS    RESTARTS   AGE
nginx     1/1       Running   0          13s
```

Now the pod is setup and ready to use.

### Get detailed information about the pod

```
$ kubectl describe pod nginx
Name:				nginx
Namespace:			default
Image(s):			nginx
Node:				127.0.0.1/127.0.0.1
Start Time:			Sun, 22 Nov 2015 23:40:10 +0000
Labels:				<none>
Status:				Running
Reason:				
Message:			
IP:				172.17.0.27
Replication Controllers:	<none>
Containers:
 nginx:
   Container ID:	docker://c2d94f95c1ac346e25386b908f4c251d2083084cbdb4d3b92a7c041672ad4521
   Image:		nginx
   Image ID:		docker://9fab4090484a840de49347c9c49597ab32df23ec26bb98d7a7ec24d59dff8945
   QoS Tier:
     cpu:		BestEffort
     memory:		BestEffort
   State:		Running
     Started:		Sun, 22 Nov 2015 23:40:22 +0000
   Ready:		True
   Restart Count:	0
   Environment Variables:
Conditions:
 Type		Status
 Ready 	True 
No volumes.
Events:
 FirstSeen	LastSeen	Count	From			SubobjectPath		Reason		Message
 ─────────	────────	─────	────			─────────────		──────		───────
 2m		2m		1	{scheduler }					Scheduled	Successfully assigned nginx to 127.0.0.1
 2m		2m		1	{kubelet 127.0.0.1}				FailedSync	Error syncing pod, skipping: DNS ResolvConfPath specified but does not exist. It could not be updated: /mnt/sda1/var/lib/docker/containers/c5225d292cdcfac66367338b1f8d35fd2bccd007f37bd1b4e904b16a87bc384d/resolv.conf
 2m		2m		1	{kubelet 127.0.0.1}	spec.containers{nginx}	Pulled		Container image "nginx" already present on machine
 2m		2m		1	{kubelet 127.0.0.1}	spec.containers{nginx}	Created		Created container with docker id c2d94f95c1ac
 2m		2m		1	{kubelet 127.0.0.1}	spec.containers{nginx}	Started		Started container with docker id c2d94f95c1ac
```

Describe gives us detailed information about the pod including its current state, which node it is running on, and lifecycle events.
We don't need this information now, but can use it if we need to debug issues with a pod.

### Get the pods container logs

```
$ kubectl logs nginx
I1122 23:48:27.326848   20553 logs.go:40] http: WriteHeader called with both Transfer-Encoding of "chunked" and a Content-Length of 0
```

The logs command fetches the STDOUT from the container running in the nginx pod.  In a multinode cluster, this coudl be running on
any node, and the logs command will find it.

### Deleting a pod

```
$ kubectl delete pod nginx
pod "nginx" deleted
```

The delete command stops the pod from running, and frees up resources on the node it was scheduled on.

### TODO: Mountain a volume in a pod - why is this useful

### TODO: Schedule multiple containers in a pod sharing the volume - why is this useful

### TODO: Reserving resources in a pod


## Replication Controllers

### Scheduling work in a *Replication Control*

```
$ kubectl run nginx --image=nginx 
```

This will schedule a single *Pod* to be run, and monitored by a *Replication Controller*.  The replication controller
will monitor the state of the pod, and create the *Pod* if it dies.  The *Pod* maybe rescheduled on a different
*Node* if necessary.

### TODO: Let the Replication Controller reschedule a pod by kill the container.  What it come back up


### TODO: Schedule multiple instances of a pod using a Replication Controller


### Services

### Create a service to load balance traffic across multiple pods



[![Analytics](https://kubernetes-site.appspot.com/UA-36037335-10/GitHub/guide/GettingStarted.md?pixel)]()
