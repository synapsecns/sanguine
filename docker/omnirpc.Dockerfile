FROM alpine

RUN apk add --no-cache ca-certificates
RUN apk add --no-cache gcc musl-dev linux-headers git libc6-compat


WORKDIR /app
COPY omnirpc /app/omnirpc

ENTRYPOINT ["/app/omnirpc"]
