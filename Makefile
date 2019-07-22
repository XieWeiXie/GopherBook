TARGET=GopherBook

default:
	@echo "Hello Golang"

vet:
	@go vet $(go list ./... | grep -v vendor)

fmt:
	@go fmt $(go list ./... | grep -v vendor)

run:
	@go run main.go

install:
	@go mod vendor -v

.PHONY default vet fmt run install
