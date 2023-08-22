FROM ghcr.io/synapsecns/sanguine-goreleaser:latest as builder

ARG VERSION=v0.0.0

# build op-proposer with the shared go.mod & go.sum files
COPY ./services /app/services
COPY ./agents /app/agents
COPY ./core /app/core
COPY ./ethergo /app/ethergo
COPY ./tools /app/tools
COPY ./contrib /app/contrib
COPY ./go.work /app/go.work
COPY ./go.work.sum /app/go.work.sum
COPY ./.git /app/.git

WORKDIR /app/services/omnirpc

RUN GOPROXY=https://proxy.golang.org go mod download

RUN CC=gcc CXX=g++  go build -tags=netgo,osusergo -ldflags="-s -w -extldflags '-static'" -o /app/bin/omnirpc  main.go

FROM alpine:3.16

COPY --from=builder /app/bin/omnirpc /usr/local/bin

CMD ["omnirpc"]
