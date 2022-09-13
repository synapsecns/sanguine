#!/bin/bash

# Protobuf Definitions

# assumes that we're only running this from make
ROOT=$(pwd)

PROTO="$ROOT/grpc/proto"

function main() {
  protoc --proto_path=grpc/proto/ --go_out=grpc/types --go-grpc_out=grpc/types  --go-grpc_opt=paths=source_relative --go_opt=paths=source_relative \
    grpc/proto/scribe/log.proto grpc/proto/scribe/types.proto grpc/proto/scribe/filter.proto grpc/proto/scribe/service.proto # add new paths here
}

main
