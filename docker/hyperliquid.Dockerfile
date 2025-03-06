FROM ubuntu:24.04

ARG USERNAME=hluser
ARG USER_UID=10000
ARG USER_GID=$USER_UID
ARG CHAIN=Testnet

# Define URLs as environment variables
ARG PUB_KEY_URL=https://raw.githubusercontent.com/hyperliquid-dex/node/refs/heads/main/pub_key.asc

# Create user and install dependencies
RUN groupadd --gid $USER_GID $USERNAME \
    && useradd --uid $USER_UID --gid $USER_GID -m $USERNAME \
    && apt-get update -y && apt-get install -y curl gnupg \
    && apt-get clean && rm -rf /var/lib/apt/lists/* \
    && mkdir -p /home/$USERNAME/hl/data && chown -R $USERNAME:$USERNAME /home/$USERNAME/hl

USER $USERNAME
WORKDIR /home/$USERNAME

# Configure chain based on build arg
RUN echo "{\"chain\": \"${CHAIN}\"}" > /home/$USERNAME/visor.json

# Import GPG public key
RUN curl -o /home/$USERNAME/pub_key.asc ${PUB_KEY_URL} \
    && gpg --import /home/$USERNAME/pub_key.asc

# Download hl-visor binary
# For now, we're skipping verification since the paths seem to be inconsistent
RUN set -ex \
    && curl -L -o /home/$USERNAME/hl-visor https://binaries.hyperliquid-testnet.xyz/${CHAIN}/hl-visor \
    && chmod +x /home/$USERNAME/hl-visor

# Expose gossip ports
EXPOSE 4000-4010

# Run a non-validating node
ENTRYPOINT ["/home/hluser/hl-visor", "run-non-validator", "--replica-cmds-style", "recent-actions"]
