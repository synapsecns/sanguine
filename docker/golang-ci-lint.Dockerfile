# USN-8944-1 removes trust in the compromised eMudhra CA chain.
FROM ubuntu:24.04 AS certificates
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && dpkg --compare-versions "$(dpkg-query -W -f='${Version}' ca-certificates)" ge 20260816ubuntu0.24.04.1 \
    && rm -rf /var/lib/apt/lists/*

FROM gcr.io/distroless/static:latest

# The base image may still carry a stale trust bundle.
COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

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
