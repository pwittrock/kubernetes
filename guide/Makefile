
all:
#	cd ../ && ./hack/build-cross.sh
	cp ../_output/dockerized/bin/linux/amd64/kubectl ./amd64/
	cp -r ../cluster/ ./cluster
#	cp ../_output/dockerized/bin/linux/amd64/kube-apiserver ./amd64/
#	cp ../third_party/etcd/etcd ./amd64/
#	cp ../third_party/etcd/etcdctl ./amd64/
	docker build -t pwittrock/k8sguide .

clean:
	rm amd64/*

.PHONY: all
