HOSTNAME=registry.terraform.io
NAMESPACE=rlmartin
VERSION := $(shell git describe --tags --dirty | sed -e 's/^v//g')
NAME=salesforce
BINARY=terraform-provider-${NAME}
OS_ARCH=darwin_arm64

.PHONY: docs

default: install

docs:
	rm -rf docs
	mkdir docs
	go generate ./...
	rm docs/*.gtpl
	rm -rf docs/client

clean:
	rm -Rf client/
	rm -Rf data_profile/
	rm -Rf models/
	rm -Rf salesforce/

generate: clean
	swagger generate client -f ./spec_files/current.json --template-dir templates -C config.yml > swagrun.log

build: generate
	go build -ldflags="-X main.version=$(VERSION) -X main.commit=n/a" -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/${BINARY}_v${VERSION}

fmt:
	gofmt -w .

fmt-check:
	gofmt .

nogen: 
	go build -ldflags="-X main.version=$(base_version) -X main.commit=n/a" -o ${BINARY}
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}/${BINARY}_v${VERSION}

test:
	TF_ACC=1 go test ./... -timeout 120m -v
