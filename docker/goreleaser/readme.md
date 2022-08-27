# Go Releaser

Go releaser is used to perform releases. We have a submodule for [goreleaser-cgo-cross-compiler](https://github.com/iotaledger/goreleaser-cgo-cross-compiler) because we need to pass in a custom build-arg for goreleaser pro (required for [monorepo](https://goreleaser.com/customization/monorepo/)). We build the base image whenever changes are present
