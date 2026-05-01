# Copyright (c) HashiCorp, Inc.
# SPDX-License-Identifier: MPL-2.0

default: fmt lint install generate

build:
	go build ./...

install: build
	go install .

lint:
	golangci-lint run

generate:
	cd tools; go generate ./...

fmt:
	gofmt -s -w -e .

test:
	go test -v -cover -timeout=120s -parallel=10 ./...

testacc:
	TF_ACC=1 go test -v -cover -timeout 120m ./...

.PHONY: build install lint generate fmt test testacc