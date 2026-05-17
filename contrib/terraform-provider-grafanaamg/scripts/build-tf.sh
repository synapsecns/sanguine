#!/bin/bash
# shellcheck disable=SC2086

TF_PREFIX=$(go version | awk '{print $NF}' | sed 's/\//_/')
PLUGIN_DIR=$(realpath -m ~/.terraform.d/plugins/registry.terraform.io/synapsecns/grafanaamg/1.0.0/$TF_PREFIX)

GOWORK=off go build .

rm -rf "$PLUGIN_DIR"
mkdir -p "$PLUGIN_DIR"
cp terraform-provider-grafanaamg "$PLUGIN_DIR"
