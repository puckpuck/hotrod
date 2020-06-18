

# The binary to build (just the basename).
BIN := hotrod

# Where to push the docker image.
REGISTRY ?= puckpuck

# Which architecture to build - see $(ALL_ARCH) for options.
ARCH ?= amd64

# Which OS to build
GOOS ?= linux

VERSION := $(shell cat version.txt | tr -d '\n')
IMAGE := $(REGISTRY)/$(BIN)

all: build

build:
	CGO_ENABLED=0 ARCH=$(ARCH) GOOS=$(GOOS) go build -o bin/$(BIN) -ldflags "-X version.VERSION=$(VERSION)" .

image:
	docker build -t $(IMAGE):$(VERSION) .

push:
	docker push $(IMAGE):$(VERSION)

release: build image push

clean:
	rm -rf bin
