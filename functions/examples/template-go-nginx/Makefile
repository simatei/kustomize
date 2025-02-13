# Copyright 2019 The Kubernetes Authors.
# SPDX-License-Identifier: Apache-2.0

.PHONY: generate license fix vet fmt test build tidy image

GOBIN = $(shell go env GOBIN)
ifeq ($(GOBIN),)
GOBIN = $(shell go env GOPATH)/bin
endif

build:
	(cd image && go build -v -o $(GOBIN)/config-function .)

all: generate license build fix vet fmt test lint tidy

fix:
	(cd image && go fix ./...)

fmt:
	(cd image && go fmt ./...)

generate:
	(which $(GOBIN)/mdtogo || go get sigs.k8s.io/kustomize/cmd/mdtogo)
	(cd image && GOBIN=$(GOBIN) go generate ./...)

license:
	(which $(GOBIN)/addlicense || go get github.com/google/addlicense)
	$(GOBIN)/addlicense  -y 2022 -c "The Kubernetes Authors." -f LICENSE_TEMPLATE .

tidy:
	(cd image && go mod tidy)

lint:
	(which $(GOBIN)/golangci-lint || go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.19.1)
	(cd image && $(GOBIN)/golangci-lint run ./...)

test:
	(cd image && go test -cover ./...)

vet:
	(cd image && go vet ./...)

image:
	docker build image -t gcr.io/kustomize-functions/example-nginx:v0.2.0
	docker push gcr.io/kustomize-functions/example-nginx:v0.2.0
