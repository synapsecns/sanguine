FROM --platform=linux/amd64 debian:11

ARG VERSION_ARG=1.20.0-pro
ARG SHA_ARG=d2d76cf4b212f67cb9995c8539167a1c6d771859aad20ed242bfab640b6d396f
ARG FILE_ARG=goreleaser-pro_Linux_x86_64.tar.gz
ARG DOWNLOAD_ARG=https://github.com/goreleaser/goreleaser-pro/releases/download/v${VERSION_ARG}/${FILE_ARG}

# from https://github.com/iotaledger/goreleaser-cgo-cross-compiler
ENV GORELEASER_VERSION=$VERSION_ARG
ENV GORELEASER_SHA=$SHA_ARG
ENV GORELEASER_DOWNLOAD_FILE=$FILE_ARG
ENV GORELEASER_DOWNLOAD_URL=$DOWNLOAD_ARG

# Golang
ENV GOLANG_VERSION=1.21.3
ENV GOLANG_SHA=1241381b2843fae5a9707eec1f8fb2ef94d827990582c7c7c32f5bdfbfd420c8
ENV GOLANG_DOWNLOAD_FILE=go${GOLANG_VERSION}.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_URL=https://dl.google.com/go/${GOLANG_DOWNLOAD_FILE}

ENV SNAP_VERSION=v4.30.6
ENV SNAP_BINARY=yq_linux_amd64


# Docker
ENV DOCKER_VERSION=24.0.5
ENV DOCKER_SHA=0a5f3157ce25532c5c1261a97acf3b25065cfe25940ef491fa01d5bea18ddc86
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
LABEL org.opencontainers.image.description="Goreleaser Docker image"

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

# Install node 20
RUN apt-get install -y curl ca-certificates gnupg zip

RUN mkdir -p /etc/apt/keyrings

RUN curl -fsSL https://deb.nodesource.com/gpgkey/nodesource-repo.gpg.key |  gpg --dearmor -o /etc/apt/keyrings/nodesource.gpg

ENV NODE_MAJOR=20

RUN echo "deb [signed-by=/etc/apt/keyrings/nodesource.gpg] https://deb.nodesource.com/node_$NODE_MAJOR.x nodistro main" |  tee /etc/apt/sources.list.d/nodesource.list

RUN apt-get update && apt-get install -y nodejs

# Install syft

RUN curl -sSfL https://raw.githubusercontent.com/anchore/syft/main/install.sh | sh -s -- -b /usr/local/bin
