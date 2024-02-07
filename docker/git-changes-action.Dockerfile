FROM alpine:latest as builder

RUN apk add --no-cache ca-certificates
RUN update-ca-certificates

# add a user here because addgroup and adduser are not available in scratch
RUN addgroup -S gitchanges \
    && adduser -S -u 10000 -g gitchanges gitchanges


FROM scratch

LABEL org.label-schema.description="Release Copier Action Docker Image"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/contrib/git-changes-action"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.description="Git Changes Action Docker image"

# copy ca certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# copy users from builder
COPY --from=builder /etc/passwd /etc/passwd

WORKDIR /git-changes-action
COPY git-changes-action /app/git-changes-action

ENTRYPOINT ["/app/git-changes-action"]
