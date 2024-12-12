FROM ghcr.io/synapsecns/sanguine-goreleaser:latest as builder

WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/root/go/pkg/mod \
    cd contrib/golang-ci-lint && \
    go build -o /app/bin/golang-ci-lint

FROM alpine:3.18
COPY --from=builder /app/bin/golang-ci-lint /usr/local/bin/
RUN chmod +x /usr/local/bin/golang-ci-lint

ENTRYPOINT ["/usr/local/bin/golang-ci-lint"]
