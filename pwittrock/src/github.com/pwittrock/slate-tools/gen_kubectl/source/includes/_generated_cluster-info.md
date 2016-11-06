# cluster-info

Display addresses of the master and services with label kubernetes.io/cluster-service=true
To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.

### Usage

`$ cluster-info`



### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
include-extended-apis |  | true | If true, include definitions of new APIs via calls to the API server. [default true] 


## dump


Dumps cluster info out suitable for debugging and diagnosing cluster problems.  By default, dumps everything to
stdout. You can optionally specify a directory with --output-directory.  If you specify a directory, kubernetes will
build a set of files in that directory.  By default only dumps things in the 'kube-system' namespace, but you can
switch to a different namespace with the --namespaces flag, or specify --all-namespaces to dump all namespaces.

The command also dumps the logs of all of the pods in the cluster, these logs are dumped into different directories
based on namespace and pod name.


### Usage

`$ dump`

> Dump current cluster state to stdout

```shell
kubectl cluster-info dump
```

> Dump current cluster state to /path/to/cluster-state

```shell
kubectl cluster-info dump --output-directory=/path/to/cluster-state
```

> Dump all namespaces to stdout

```shell
kubectl cluster-info dump --all-namespaces
```

> Dump a set of namespaces to /path/to/cluster-state

```shell
kubectl cluster-info dump --namespaces default,kube-system --output-directory=/path/to/cluster-state
```


### Flags

Name | Shorthand | Default | Usage
---- | --------- | ------- | ----- 
all-namespaces |  | false | If true, dump all namespaces.  If true, --namespaces is ignored. 
namespaces |  | [] | A comma separated list of namespaces to dump. 
output-directory |  |  | Where to output the files.  If empty or '-' uses stdout, otherwise creates a directory hierarchy in that directory 


