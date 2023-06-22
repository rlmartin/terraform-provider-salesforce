HOSTNAME=terraform.local
NAMESPACE=com
NAME=salesforce
BINARY=terraform-provider-${NAME}
VERSION=0.2.0
OS_ARCH=darwin_arm64

default: install

clean:
	rm -Rf client/
	rm -Rf data_profile/
	rm -Rf models/
	rm -Rf salesforce/

build:  clean
	swagger generate client -f ./spec_files/current.json --template-dir templates -C config.yml > swagrun.log
	go build -o ${BINARY}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	cp ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

nogen: 
	go build -o ${BINARY}
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}