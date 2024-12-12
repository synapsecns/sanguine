# syntax=docker/dockerfile:1.4

FROM ghcr.io/synapsecns/sanguine-goreleaser:latest as builder

WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/root/go/pkg/mod \
    cd contrib/golang-ci-lint && \
    go build -o /app/bin/golang-ci-lint

FROM alpine:3.18

LABEL org.label-schema.description="Golang-CI-Lint Docker Image"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/golang-ci-lint"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="Golang-CI-Lint Docker image"

COPY --from=builder /app/bin/golang-ci-lint /usr/local/bin/
RUN chmod +x /usr/local/bin/golang-ci-lint

ENTRYPOINT ["/usr/local/bin/golang-ci-lint"]
