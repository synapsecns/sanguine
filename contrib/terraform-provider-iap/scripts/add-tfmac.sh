#!/bin/zsh

# if not already present in zshrc
if [ "$(grep -c -w "alias tfmac='TFENV_ARCH=arm64 TFENV_TERRAFORM_VERSION=latest:^1.3 terraform'" ~/.zshrc)" -le 0 ]; then
    echo "adding tfmac command to zshrc. You might have to source ~/.zshrc or open a new tab"
    echo "alias tfmac='TFENV_ARCH=arm64 TFENV_TERRAFORM_VERSION=latest:^1.3 terraform'" >> ~/.zshrc
fi
