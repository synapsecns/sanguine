FROM alpine as builder

RUN apk add --no-cache ca-certificates

FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

WORKDIR /app
COPY omnirpc /app/omnirpc

ENTRYPOINT ["/app/omnirpc"]
