FROM alpine:latest

WORKDIR /app
COPY main /app/main

ENTRYPOINT ["/app/main"]
