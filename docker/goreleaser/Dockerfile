FROM gohornet/goreleaser-cgo-cross-compiler:1.19

LABEL org.label-schema.description="GoReleaser synapse node dockerfile"
LABEL org.label-schema.name="ghcr.io/synapsecns/synapse-node-goreleaser"
LABEL org.label-schema.schema-version="1.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/synapse-node/tree/deploydoc/docker/goreleaser"
LABEL org.label-schema.usage="https://github.com/synapsecns/synapse-node/blob/deploydoc/docker/goreleaser/readme.md"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/synapse-node"

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
CMD ["goreleaser", "-v"]
