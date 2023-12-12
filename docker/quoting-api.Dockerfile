FROM gcr.io/distroless/static:latest

LABEL org.label-schema.description="Quoting API Docker file"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/quoting-api"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="Quoting API Docker image"

USER nonroot:nonroot

WORKDIR /app
COPY --chown=nonroot:nonroot quoting-api /app/quoting-api

ENTRYPOINT ["/app/quoting-api"]
