.PHONY: all get build

get:
	go get -v -d ./...

build:
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o bin/quotas cmd/quotas/*.go
	chmod +x bin/quotas

default: get build