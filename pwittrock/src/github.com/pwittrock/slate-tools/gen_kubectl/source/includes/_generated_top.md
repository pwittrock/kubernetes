# top


Display Resource (CPU/Memory/Storage) usage.

The top command allows you to see the resource consumption for nodes or pods.

### Usage

`$ top`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 


## node


Display Resource (CPU/Memory/Storage) usage of nodes.

The top-node command allows you to see the resource consumption of nodes.

### Usage

`$ node [NAME | -l label]`

> Show metrics for all nodes

```shell
kubectl top node
```

> Show metrics for a given node

```shell
kubectl top node NODE_NAME
```


### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
selector | l |  | Selector (label query) to filter on 


## pod


Display Resource (CPU/Memory/Storage) usage of pods.

The 'top pod' command allows you to see the resource consumption of pods.

Due to the metrics pipeline delay, they may be unavailable for a few minutes
since pod creation.

### Usage

`$ pod [NAME | -l label]`

> Show metrics for all pods in the default namespace

```shell
kubectl top pod
```

> Show metrics for all pods in the given namespace

```shell
kubectl top pod --namespace=NAMESPACE
```

> Show metrics for a given pod and its containers

```shell
kubectl top pod POD_NAME --containers
```

> Show metrics for the pods defined by label name=myLabel

```shell
kubectl top pod -l name=myLabel
```


### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
all-namespaces |  | false | If present, list the requested object(s) across all namespaces. Namespace in current context is ignored even if specified with --namespace. 
containers |  | false | If present, print usage of containers within a pod. 
selector | l |  | Selector (label query) to filter on 


