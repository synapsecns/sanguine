FROM gcr.io/distroless/static:latest

LABEL org.label-schema.description="Sin-Executor Dockerfile"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/sin-executor"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="Sin-Executor Docker image"

USER nonroot:nonroot

WORKDIR /app
COPY --chown=nonroot:nonroot sin-executor /app/sin-executor

ENTRYPOINT ["/app/sin-executor"]
