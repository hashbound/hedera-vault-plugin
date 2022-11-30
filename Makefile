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

test:
	vault write hedera/keys algo="ED25519" id="1"
	vault read hedera/keys id="1"
	vault write hedera/keys algo="ECDSA" curve="secp256k1" id="2"
	vault read hedera/keys id="2"
	vault list hedera/keys
	vault write hedera/keys/import id="3" algo="ED25519" privateKey="302e020100300506032b65700422042091132178e72057a1d7528025956fe39b0b847f200ab59b2fdd367017f3087137"
	vault read hedera/keys id="3"
	vault write hedera/keys/2/sign message="123"
	vault delete hedera/keys id="1"

	vault write hedera/accounts/import id="1" keyId="3" accountId="0.0.2"
	vault read hedera/accounts id="1"
	vault list hedera/accounts
	vault write hedera/accounts/sign_transaction id="1" transaction="0a742a720a6e0a120a0c0888e09d9c0610ddccbb8c0212021802120218031880bcc1960b22020878ea01490a0b4a696d6d7920546f6b656e12024a54180220e8072a021802322212200aa8e21064c61eab86e2a9c164565b4e7a9a4146106e0a6cd03a8c395a110e92720218027a0508d0c8e1031200"

.PHONY: build clean fmt start enable