FROM debian:11

ARG VERSION_ARG=1.13.1
ARG SHA_ARG=136fecfb2e2f3a7965274ad5e2571985d8b2fa724b6536874f082e4b0bb9f344
ARG FILE_ARG=goreleaser_Linux_x86_64.tar.gz
ARG DOWNLOAD_ARG=https://github.com/goreleaser/goreleaser/releases/download/v${VERSION_ARG}/${FILE_ARG}

# from https://github.com/iotaledger/goreleaser-cgo-cross-compiler
ENV GORELEASER_VERSION=$VERSION_ARG
ENV GORELEASER_SHA=$SHA_ARG
ENV GORELEASER_DOWNLOAD_FILE=$FILE_ARG
ENV GORELEASER_DOWNLOAD_URL=$DOWNLOAD_ARG

# Golang
ENV GOLANG_VERSION=1.19
ENV GOLANG_SHA=464b6b66591f6cf055bc5df90a9750bf5fbc9d038722bb84a9d56a2bea974be6
ENV GOLANG_DOWNLOAD_FILE=go${GOLANG_VERSION}.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_URL=https://dl.google.com/go/${GOLANG_DOWNLOAD_FILE}

ENV SNAP_VERSION=v4.30.6
ENV SNAP_BINARY=yq_linux_amd64


# Docker
ENV DOCKER_VERSION=20.10.9
ENV DOCKER_SHA=caf74e54b58c0b38bb4d96c8f87665f29b684371c9a325562a3904b8c389995e
ENV DOCKER_DOWNLOAD_FILE=docker-${DOCKER_VERSION}.tgz
ENV DOCKER_DOWNLOAD_URL=https://download.docker.com/linux/static/stable/x86_64/${DOCKER_DOWNLOAD_FILE}

# Install cross compiling tools
RUN apt-get update && apt-get install -y build-essential wget git cmake \
	libc6-dev-armel-cross libc6-dev-armel-cross binutils-arm-linux-gnueabi libncurses5-dev \
	gcc-mingw-w64 g++-mingw-w64 \
	gcc-aarch64-linux-gnu g++-aarch64-linux-gnu && \
	apt-get -y autoremove


# Download yq
# TODO: sha256sum
RUN wget https://github.com/mikefarah/yq/releases/download/${SNAP_VERSION}/${SNAP_BINARY} -O /usr/bin/yq \
    && chmod +x /usr/bin/yq

# Download Docker
RUN	wget ${DOCKER_DOWNLOAD_URL}; \
	echo "${DOCKER_SHA} ${DOCKER_DOWNLOAD_FILE}" | sha256sum -c - || exit 1; \
	tar --extract --file ${DOCKER_DOWNLOAD_FILE} --strip-components 1 --directory /usr/local/bin/; \
	rm ${DOCKER_DOWNLOAD_FILE}

# Download GoReleaser
RUN wget ${GORELEASER_DOWNLOAD_URL}; \
	echo "${GORELEASER_SHA} ${GORELEASER_DOWNLOAD_FILE}" | sha256sum -c - || exit 1; \
	tar -xzf ${GORELEASER_DOWNLOAD_FILE} -C /usr/bin/ goreleaser; \
	rm ${GORELEASER_DOWNLOAD_FILE};

# Download Golang
RUN wget ${GOLANG_DOWNLOAD_URL}; \
	echo "${GOLANG_SHA} ${GOLANG_DOWNLOAD_FILE}" | sha256sum -c - || exit 1; \
	tar -xzf ${GOLANG_DOWNLOAD_FILE} -C /usr/local; \
	rm ${GOLANG_DOWNLOAD_FILE};


# Add Golang to PATH
ENV PATH=${PATH}:/usr/local/go/bin

CMD ["goreleaser", "-v"]

# end copied section

LABEL org.label-schema.description="GoReleaser sanguine dockerfile"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine-goreleaser"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"

# add any extra dependencies here

# see: https://github.com/tendermint/tm-db/blob/8f92601b6539a3611063709f65256150e198cfd9/tools/Dockerfile
ENV LD_LIBRARY_PATH=/usr/local/lib


RUN apt-get update && apt-get install -y --no-install-recommends \
    libbz2-dev libgflags-dev libsnappy-dev libzstd-dev zlib1g-dev \
    make tar wget

ARG ROCKSDB=6.24.2
ENV ROCKSDB_DISABLE_SNAPPY=1
ENV ROCKSDB_DISABLE_GFLAGS=1
ENV ROCKSDB_DISABLE_ZSTD=1


RUN apt-get update
RUN \
  wget -q https://github.com/facebook/rocksdb/archive/v${ROCKSDB}.tar.gz \
  && tar -zxf v${ROCKSDB}.tar.gz \
  && cd rocksdb-${ROCKSDB} \
  && DEBUG_LEVEL=0 make static_lib \
  && make install \
  && ldconfig \
  && cd .. \
  && rm -rf v${ROCKSDB}.tar.gz rocksdb-${ROCKSDB}


ENV PATH=${PATH}:/etc/musl/x86_64-linux-musl-native/bin:/etc/musl/aarch64-linux-musl-cross/bin:/usr/local/go/bin

# Install node 16
RUN apt-get install -y curl

RUN curl -fsSL https://deb.nodesource.com/setup_16.x | bash -

RUN apt-get install -y nodejs
