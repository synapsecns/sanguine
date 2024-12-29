FROM gcr.io/distroless/static:latest

LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="Golang-CI-Lint version manager for Sanguine"
LABEL org.opencontainers.image.name="ghcr.io/synapsecns/sanguine/golang-ci-lint"
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.vendor="Synapse Labs"
LABEL org.opencontainers.image.documentation="https://github.com/synapsecns/sanguine/tree/master/contrib/golang-ci-lint"

USER nonroot:nonroot

WORKDIR /app
COPY --chown=nonroot:nonroot golang-ci-lint /app/golang-ci-lint

ENTRYPOINT ["/app/golang-ci-lint"]
