.PHONY: all

fmt:
	go fmt ./...

vet:
	go vet ./...

build: fmt vet