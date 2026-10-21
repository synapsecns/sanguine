# use a custom image to add labels, please use features instead of modifying this directly
FROM mcr.microsoft.com/devcontainers/base:ubuntu-24.04


LABEL org.label-schema.description="Sanguin Devcontainer Docker Image"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/.devcontainer"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="Sanguine dev container"

# Require the Ubuntu certificate security update before HTTPS downloads.
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && dpkg --compare-versions "$(dpkg-query -W -f='${Version}' ca-certificates)" ge 20260816ubuntu0.24.04.1 \
    && rm -rf /var/lib/apt/lists/*

# TODO: use a feature
RUN curl -L https://foundry.paradigm.xyz | bash
# see: https://stackoverflow.com/a/25423366
SHELL ["/bin/bash", "-c"]
RUN ~/.foundry/bin/foundryup
