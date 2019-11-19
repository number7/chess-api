ROOT_DIR := $(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
GOPATH := $(shell go env GOPATH)

GOOS ?= $(shell uname | tr 'A-Z' 'a-z')
GOARCH := amd64

executable = $(notdir $(shell pwd))

build: ## Build binary.
	@echo "building ${executable}"
	GOOS=${GOOS} GOARCH=${GOARCH} go build -o ./bin/${executable} ./cmd/${executable}
