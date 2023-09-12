# Go Releaser
<!-- see: https://ghcr-badge.egpl.dev/ -->
![Image Size](https://ghcr-badge.egpl.dev/synapsecns/sanguine-goreleaser/size?color=%2344cc11&tag=latest&label=image+size&trim=)

Go releaser is used to perform releases. We have a submodule for [goreleaser-cgo-cross-compiler](https://github.com/iotaledger/goreleaser-cgo-cross-compiler) because we need to pass in a custom build-arg for goreleaser pro (required for [monorepo](https://goreleaser.com/customization/monorepo/)). We build the base image whenever changes are present

Use: `docker pull ghcr.io/synapsecns/sanguine-goreleaser:latest`



### Supported OS and architectures:

- Windows (amd64)
- Linux (amd64, ARM64)
- macOS (amd64) **No CGO support**

### Used versions

- **GoLang**: 1.21.0
- **GoReleaser Pro**: 1.19.2
- **Docker**: 24.0.5
