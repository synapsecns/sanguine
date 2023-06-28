FROM gcr.io/distroless/static:latest

LABEL org.label-schema.description="CCTP Relayer Dockerfile"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/scribe"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="CCTP Relayer Docker image"

USER nonroot:nonroot

WORKDIR /app
COPY --chown=nonroot:nonroot cctp-relayer /app/cctp-relayer

ENTRYPOINT ["/app/cctp-relayer"]
