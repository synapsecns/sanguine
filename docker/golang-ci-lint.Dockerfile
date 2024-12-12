FROM alpine:3.18

COPY golang-ci-lint /usr/local/bin/
RUN chmod +x /usr/local/bin/golang-ci-lint

ENTRYPOINT ["/usr/local/bin/golang-ci-lint"]
