# USN-8944-1 removes trust in the compromised eMudhra CA chain.
FROM ubuntu:24.04 AS certificates
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && dpkg --compare-versions "$(dpkg-query -W -f='${Version}' ca-certificates)" ge 20260816ubuntu0.24.04.1 \
    && rm -rf /var/lib/apt/lists/*

FROM gcr.io/distroless/static:latest

# The base image may still carry a stale trust bundle.
COPY --from=certificates /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

LABEL org.label-schema.description="CCTP Relayer Dockerfile"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/services/cctp-relayer"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="CCTP Relayer Docker image"

USER nonroot:nonroot

WORKDIR /app
COPY --chown=nonroot:nonroot cctp-relayer /app/cctp-relayer

ENTRYPOINT ["/app/cctp-relayer"]
