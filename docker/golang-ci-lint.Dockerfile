# syntax=docker/dockerfile:1.4

FROM golang:1.21-alpine AS builder

WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o /build/bin/golang-ci-lint \
    -ldflags="-s -w -extldflags '-static'" \
    -tags netgo,osusergo \
    -trimpath

FROM alpine:3.18

LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="Golang-CI-Lint version manager for Sanguine"
LABEL org.opencontainers.image.name="ghcr.io/synapsecns/sanguine/golang-ci-lint"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.vendor="Synapse Labs"
LABEL org.opencontainers.image.documentation="https://github.com/synapsecns/sanguine/tree/master/contrib/golang-ci-lint"

COPY --from=builder /build/bin/golang-ci-lint /usr/local/bin/
RUN chmod +x /usr/local/bin/golang-ci-lint

ENTRYPOINT ["/usr/local/bin/golang-ci-lint"]
