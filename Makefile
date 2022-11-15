GOARCH = amd64

UNAME = $(shell uname -s)

ifndef OS
	ifeq ($(UNAME), Linux)
		OS = linux
	else ifeq ($(UNAME), Darwin)
		OS = darwin
	endif
endif

all: build start

build:
	GOOS=$(OS) GOARCH="$(GOARCH)" go build -o build/hedera-vault-plugin ./main.go

start:
	vault server -dev -dev-root-token-id=root -dev-plugin-dir=./build

enable:
	vault secrets enable -path=hedera hedera-vault-plugin

clean:
	rm -f ./build/hedera-vault-plugin

fmt:
	go fmt $$(go list ./...)

.PHONY: build clean fmt start enable