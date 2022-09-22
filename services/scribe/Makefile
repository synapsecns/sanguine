include ../../make/go.Makefile

BUF_VERSION := "1.7.0"
BIN="/usr/local/bin"

setup-go-path:
	@eval $(export PATH="$PATH:$(go env GOPATH)/bin")

protoc-install:
	@#Github Actions
	@if [ "$(shell which protoc)" = "" ] && [ "$(GITHUB_WORKFLOW)" != "" ]; then sudo apt install -y protobuf-compiler; fi;
	@#Brew - MacOS
	@if [ "$(shell which protoc)" = "" ] && [ "$(shell which brew)" != "" ] && [ "$(GITHUB_WORKFLOW)" == "" ]; then brew install protobuf; fi;


protoc-gen-install: setup-go-path
	@if [ "$(shell which protoc-gen-go)" = "" ]; then go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1; fi;

	@if [ "$(shell which protoc-gen-go)" = "" ]; then echo "could not find protoc-gen-go even after installing, you might need to add your gopath to bash. You can also rerun the command in a new tab";  fi;

	@if [ "$(shell protoc-gen-go --version)" != "protoc-gen-go v1.28.1" ]; then echo "Wrong version of protoc-gen-go installed, needed: v1.28.1";  fi;

protoc-gen-grpc-gateway-install: setup-go-path
	@if [ "$(shell which protoc-gen-grpc-gateway)" = "" ]; then go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.11.3; fi;

	@if [ "$(shell which protoc-gen-grpc-gateway)" = "" ]; then echo "could not find protoc-gen-grpc-gateway even after installing, you might need to add your gopath to bash. You can also rerun the command in a new tab";  fi;


grpc-install: setup-go-path
	@if [ "$(shell which protoc-gen-go-grpc)" = "" ]; then go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2; fi;

	@if [ "$(shell which protoc-gen-go-grpc)" = "" ]; then echo "could not find protoc-gen-go-grpc even after installing, you might need to add your gopath to bash. You can also rerun the command in a new tab";  fi;

	@if [ "$(shell  protoc-gen-go-grpc --version)" != "protoc-gen-go-grpc 1.2.0" ]; then echo "Wrong version of protoc-gen-go installed, needed: v1.2.0";  fi;

openapi-install: setup-go-path
	@if [ "$(shell which protoc-gen-openapiv2)" = "" ]; then go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.11.3; fi;

	@if [ "$(shell which protoc-gen-openapiv2)" = "" ]; then echo "could not find protoc-gen-openapiv2 even after installing, you might need to add your gopath to bash. You can also rerun the command in a new tab";  fi;


buf-install:
	@if [ "$(shell which buf)" = "" ]; then ./scripts/install-buf.sh $(BUF_VERSION); fi;

swagger-codegen-install:
	@#Brew - MacOS
	@if [ "$(shell which swagger-codegen)" = "" ] && [ "$(shell which brew)" != "" ]; then brew install swagger-codegen; fi;

# TODO: should be replaced w/ https://github.com/bufbuild/buf
# (see: https://github.com/grpc-ecosystem/grpc-gateway)
generate: openapi-install swagger-codegen-install buf-install protoc-gen-grpc-gateway-install protoc-install grpc-install protoc-gen-install  ## generate protobufs
	buf generate
	@eval $$(cd $(GIT_ROOT)/services/scribe); ./scripts/codegen.sh

generate-ci: generate ## CI only command for generate prereqs
