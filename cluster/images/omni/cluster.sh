#!/bin/bash

# RUN
# docker build -t pwittrock/omni .
# docker run -ti --net=host --name k8s_console -v /var/run/docker.sock:/var/run/docker.sock -v $(which docker):/bin/docker pwittrock/omni

#started=`docker inspect -f {{.State.Running}} k8s_etcd`
#if ! $started ; then
#fi

docker inspect -f {{.State.Running}} omnik8s_etcd >> /dev/null 2>&1
found=$?
if [ $found -eq 1 ]; then
echo "creating omnik8s_etcd container..."
docker run \
    --name=omnik8s_etcd \
    --net=host \
    -d \
    gcr.io/google_containers/etcd:2.0.12 \
        /usr/local/bin/etcd \
        --addr=127.0.0.1:4001 \
        --bind-addr=0.0.0.0:4001 \
        --data-dir=/var/etcd/data
echo "done"
fi
started=`docker inspect -f {{.State.Running}} omnik8s_etcd`
if ! $started ; then
echo "starting omnik8s_etcd container..."
docker start omnik8s_etcd
echo "done"
fi

docker inspect -f {{.State.Running}} omnik8s_kubelet >> /dev/null 2>&1
found=$?
if [ $found -eq 1 ]; then
echo "creating omnik8s_kubelet container..."
docker run \
    --name=k8s_kubelet \
    --volume=/:/rootfs:ro \
    --volume=/sys:/sys:ro \
    --volume=/dev:/dev \
    --volume=/var/lib/docker/:/var/lib/docker:ro \
    --volume=/var/lib/kubelet/:/var/lib/kubelet:rw \
    --volume=/var/run:/var/run:rw \
    --pid=host \
    gcr.io/google_containers/hyperkube:v1.1.2 \
        /hyperkube kubelet \
        --containerized \
        --hostname-override="127.0.0.1" \
        --hostname-override="127.0.0.1" \
        --address="0.0.0.0" \
        --api-servers=http://localhost:8080 \
        --config=/etc/kubernetes/manifests
echo "done"
fi
started=`docker inspect -f {{.State.Running}} omnik8s_kubelet`
if ! $started ; then
echo "starting omnik8s_kubelet container..."
docker start omnik8s_kubelet
echo "done"
fi

docker inspect -f {{.State.Running}} omnik8s_proxy >> /dev/null 2>&1
found=$?
if [ $found -eq 1 ]; then
echo "creation omnik8s_proxy container..."
docker run \
    --name=omnik8s_proxy \
    --net=host \
    --pid=host \
    --privileged=true \
    -d \
    gcr.io/google_containers/hyperkube:v1.1.2 \
        /hyperkube proxy \
        --master=http://127.0.0.1:8080 \
        --v=2
echo "done"
fi
started=`docker inspect -f {{.State.Running}} omnik8s_proxy`
if ! $started ; then
echo "starting omnik8s_proxy container..."
docker start omnik8s_proxy
echo "done"
fi

kubectl get nodes >> /dev/null 2>&1
apiup=$?
start=`date +%s`
while [ $apiup -eq 1 ]; do
  waited=$(( `date +%s` - start ))
  echo "Waited $waited seconds for cluster to become healthy (may take a couple minutes)..."
  sleep 5
  kubectl get nodes >> /dev/null 2>&1
  apiup=$?
done


bash

# RUN
# docker build -t pwittrock/omni .
# docker run -ti --net=host --name k8s_console -v /var/run/docker.sock:/var/run/docker.sock -v $(which docker):/bin/docker pwittrock/omni