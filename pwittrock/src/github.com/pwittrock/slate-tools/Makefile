all: main.go
	go build main.go

clean:
	rm -f main
	rm -rf $(shell pwd)/gen_kubectl/build
	rm -rf $(shell pwd)/gen_open_api/build
	rm -f third_party/slate/source/includes/*.md
	rm -f third_party/slate/source/index.html.md

cli: clean
	GOPATH=$(shell pwd)/../../../.. && go run main.go --yaml-file gen_kubectl/kubectl.yaml --doc-type kubectl --template-dir gen_kubectl/ --build-dir third_party/slate/source/
	docker run -ti -v $(shell pwd)/third_party/slate/source:/slate/source -v $(shell pwd)/gen_kubectl/build:/slate/build pwittrock/slate

// Usage: REPO="your/repo" TAG="vN" make pushcli
pushcli:
	cd gen_kubectl && docker build . -t $(REPO):$(TAG)
	docker push $(REPO):$(TAG)

api: clean
	GOPATH=$(shell pwd)/../../../.. && go run main.go --yaml-file gen_open_api/config.yaml --doc-type open-api --template-dir gen_open_api/ --build-dir third_party/slate/source/ --open-api-dir gen_open_api/openapi-spec/
	docker run -v $(shell pwd)/third_party/slate/source:/slate/source -v $(shell pwd)/gen_open_api/build:/slate/build pwittrock/slate

// Usage: REPO="your/repo" TAG="vN" make pushapi
pushapi:
	cd gen_open_api && docker build . -t $(REPO):$(TAG)
	docker push $(REPO):$(TAG)
