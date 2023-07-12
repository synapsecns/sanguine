FROM gcr.io/distroless/static:latest

LABEL org.label-schema.description="Explorer Docker file"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/explorer"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="Synapse Explorer Docker image"

USER nonroot:nonroot

WORKDIR /app
COPY --chown=nonroot:nonroot explorer /app/explorer

ENTRYPOINT ["/app/explorer"]
