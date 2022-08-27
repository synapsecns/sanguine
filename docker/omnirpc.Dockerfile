FROM alpine:latest

WORKDIR /app
COPY omnirpc /app/omnirpc

ENTRYPOINT ["/app/omnirpc"]
