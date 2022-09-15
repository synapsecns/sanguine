#!/bin/bash

swagger-codegen generate -i "$(pwd)/api/static/swagger/types/v1/service.swagger.json" -l go -o grpc/client/rest/ --additional-properties packageName=rest
rm -rf grpc/client/rest/*.sh
rm -rf grpc/client/rest/.travis.yml
go fmt ./...
