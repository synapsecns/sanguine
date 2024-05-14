#!/bin/sh

cargo install postman2openapi-cli
# As go:generate does not support unix pipes, we need to do this here
postman2openapi collection.json > openapi.yaml
