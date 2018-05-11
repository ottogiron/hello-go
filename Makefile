SOURCE_DIRS := ./...

all: test build

test: lint
	go test -v ./... -cover

lint:
	go fmt $(SOURCE_DIRS)
	go vet $(SOURCE_DIRS)

build:
	docker build -t ottogiron/training-go .