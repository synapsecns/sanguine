FROM ghcr.io/synapsecns/sanguine-goreleaser:latest as builder

ARG VERSION=v0.0.0

COPY ./services /app/services
COPY ./agents /app/agents
COPY ./core /app/core
COPY ./ethergo /app/ethergo
COPY ./tools /app/tools
COPY ./contrib /app/contrib
COPY ./go.work /app/go.work
COPY ./go.work.sum /app/go.work.sum
COPY ./.git /app/.git

WORKDIR /app/services/scribe

RUN --mount=type=cache,target=/root/go/pkg/mod GOPROXY=https://proxy.golang.org go mod download -x
RUN --mount=type=cache,target=/root/go/pkg/mod  --mount=type=cache,target=/root/.cache/go-build CC=gcc CXX=g++ go build -tags=netgo,osusergo -ldflags="-s -w -extldflags '-static'" -o /app/bin/scribe  main.go

FROM ubuntu:latest

RUN apt update && apt install -y bash sqlite3 htop
COPY --from=builder /app/bin/scribe /usr/local/bin

CMD ["scribe"]
