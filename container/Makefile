CONTAINER_NAME=verosk/dummy_server
CONTAINER_VERSION=latest

dummy_server: dummy_server.go
	CGO_ENABLED=0 GOOS=linux go build -a -tags netgo -ldflags '-w' dummy_server.go

	docker build -t dummy_server .
	docker tag dummy_server ${CONTAINER_NAME}:${CONTAINER_VERSION}
	docker push ${CONTAINER_NAME}
