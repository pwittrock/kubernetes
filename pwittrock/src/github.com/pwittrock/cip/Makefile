all: main.go dockerhub/*.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

push:
	docker build . -t pwittrock/container-image-pusher:$(TAG)
	docker push pwittrock/container-image-pusher:$(TAG)
