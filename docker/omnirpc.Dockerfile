FROM gcr.io/distroless/static:latest

LABEL org.label-schema.description="Omnirpc Docker file"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/omnirpc"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="Omnirpc Docker image"

USER nonroot:nonroot

WORKDIR /app
COPY --chown=nonroot:nonroot omnirpc /app/omnirpc

ENTRYPOINT ["/app/omnirpc"]
