#!/bin/bash
# shellcheck disable=SC2086

# This gets the arch prefix we use when building a terraform plugin
TF_PREFIX=$( go version |  awk '{print $NF}' | sed 's/\//_/')

# define the plugin directory
PLUGIN_DIR=$(realpath -m ~/.terraform.d/plugins/example-helm.com/provider/helmproxy/1.0.0/$TF_PREFIX)

# fixes async problems on arm64 https://github.com/hashicorp/terraform-provider-aws/issues/20274#issuecomment-996795241
# we don't need this for production builds, just darwinarm64.
GODEBUG=asyncpreemptoff=1  go build .

# make the plugin directory if it doesn't exist
rm -rf $PLUGIN_DIR
mkdir -p $PLUGIN_DIR
cp terraform-provider-helmproxy $PLUGIN_DIR
