SERVICE_NAME=showstart

build:*.go
	 CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -o bin/${SERVICE_NAME}
build_debug:*.go
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64  go build -gcflags "all=-N -l" -o bin/$(SERVICE_NAME)
image:
	docker build --rm .
local_image:
	docker build --rm -t showstart .
local_image_release:
	docker build --rm -t showstart .