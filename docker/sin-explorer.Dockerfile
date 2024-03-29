FROM gcr.io/distroless/static:latest

LABEL org.label-schema.description="SIN Explorer Docker Image"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/services/sin-explorer"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="SIN Explorer Docker image"

USER nonroot:nonroot

WORKDIR /app
COPY --chown=nonroot:nonroot sin-explorer /app/sin-explorer

ENTRYPOINT ["/app/sin-explorer"]
