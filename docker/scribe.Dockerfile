FROM ubuntu:latest

LABEL org.label-schema.description="Scribe Docker file"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/scribe"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"

RUN apt update && apt install -y git
RUN git clone https://github.com/torvalds/linux
#USER nonroot:nonroot

WORKDIR /app
COPY scribe /app/scribe

ENTRYPOINT ["/app/scribe"]
