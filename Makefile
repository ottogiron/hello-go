VERSION := 0.1.0
SOURCE_DIRS := ./...
IMAGE_NAME := ottogiron/training-go
all: test build

test: lint
	go test -v ./... -cover

lint:
	go fmt $(SOURCE_DIRS)
	go vet $(SOURCE_DIRS)

build:
	docker build -t $(IMAGE_NAME) .

push_latest: 
	docker push $(IMAGE_NAME)

push_release:
	docker push $(IMAGE_NAME):$(VERSION)

tag_current_version: 
	docker tag $(IMAGE_NAME) $(IMAGE_NAME):$(VERSION)

release: build tag_current_version push_latest push_release
	echo "Releasing..."
	
