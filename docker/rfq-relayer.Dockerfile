FROM gcr.io/distroless/static:latest

LABEL org.label-schema.description="RFQ Relayer Docker file"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/rfq-relayer"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="RFQ Relayer Docker image"

USER nonroot:nonroot

WORKDIR /app
COPY --chown=nonroot:nonroot rfq-relayer /app/rfq-relayer

ENTRYPOINT ["/app/rfq-relayer"]
