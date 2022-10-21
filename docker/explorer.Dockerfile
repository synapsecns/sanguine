# TODO: we should use alpine here
FROM ubuntu:latest

LABEL org.label-schema.description="Explorer Docker file"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/explorer"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"

RUN apt-get update
RUN apt-get install ca-certificates -y
RUN update-ca-certificates

WORKDIR /app
COPY explorer /app/explorer

ENTRYPOINT ["/app/explorer"]
