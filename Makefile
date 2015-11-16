MAKEFILE := $(abspath $(lastword $(MAKEFILE_LIST)))
PROJECT := $(dir $(MAKEFILE))
GOPATH := $(PROJECT)/vendor
VENDOR := $(GOPATH)/src/github.com/blakepark/mongodb-rest-api
BIN := mongodb-rest-api

export PROJECT
export GOPATH
export VENDOR
export CGO_ENABLED=0

default: build

deps:
	go get github.com/gorilla/mux
	go get github.com/gorilla/context
	go get gopkg.in/mgo.v2

build:
	@rm -rf $(VENDOR)

	@mkdir -p $(PROJECT)/bin
	@mkdir -p $(VENDOR)

	@cp -r $(PROJECT)/context $(VENDOR)/context
	@cp -r $(PROJECT)/route $(VENDOR)/route
	@cp -r $(PROJECT)/mongodb $(VENDOR)/mongodb

	@go build -a -installsuffix cgo -ldflags '-w' -o $(PROJECT)/bin/$(BIN)

run: build
	$(PROJECT)/bin/$(BIN)
