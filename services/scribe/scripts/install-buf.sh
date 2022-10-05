#!/bin/bash

BIN="/usr/local/bin"

echo "installing buf, this may require sudo"
VERSION="1.7.0" && \
  sudo curl -sSL \
    "https://github.com/bufbuild/buf/releases/download/v${VERSION}/buf-$(uname -s)-$(uname -m)" \
    -o "${BIN}/buf" && \
  sudo chmod +x "${BIN}/buf"
