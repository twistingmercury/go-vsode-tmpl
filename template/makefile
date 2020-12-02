default: build

BUILD_DATE=$(shell date +%Y.%m.%d)
BUILD_VER=0.0.1
OUT_NAME=rest_api
GIT_COMMIT=$(shell git rev-parse --short HEAD)

build: 
	rm -rf _bin/*
	go build -ldflags "-X 'main.buildDate=$(BUILD_DATE)' -X 'main.buildVersion=$(BUILD_VER)' -X 'main.buildCommit=$(GIT_COMMIT)'" -o _bin/$(OUT_NAME)

run-dev: 
	go run main.go
	
run-v: build
	_bin/rest_api -c local-config -v

run: build
	_bin/rest_api -c local-config

test:
	go test -v ./... -cover

cover:
	go test -v ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out