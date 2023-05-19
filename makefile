SERVICE_NAME=showstart

build:*.go
	CGO_ENABLED="1" go build -o bin/${SERVICE_NAME}
build_debug:*.go
	CGO_ENABLED="1" go build -gcflags "all=-N -l" -o bin/$(SERVICE_NAME)
image:
	docker build  .
local_image:
	docker build -t showstart .
local_image_release:
	docker build -t showstart .