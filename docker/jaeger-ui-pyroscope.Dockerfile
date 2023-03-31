FROM node:14.17.6-alpine3.12

RUN apk add --no-cache git

WORKDIR /opt/jaeger-ui

# RUN git clone https://github.com/pyroscope-io/jaeger-ui.git /opt/jaeger-ui && git checkout 0b4bdd6a488c0d73265578f1dcb006affb76d4bd
RUN git clone https://github.com/pyroscope-io/jaeger-ui.git /opt/jaeger-ui && git checkout c3f4fa9ef7b743cf654ca4c108c55d1ade98c6a0

LABEL org.label-schema.description="Jaeger UI Pyroscope Docker file"
LABEL org.label-schema.name="ghcr.io/synapsecns/sanguine/docker"
LABEL org.label-schema.schema-version="1.0.0"
LABEL org.label-schema.vcs-url="https://github.com/synapsecns/sanguine"
LABEL org.opencontainers.image.source="https://github.com/synapsecns/sanguine"


RUN yarn install || true
ENV HOST=0.0.0.0
RUN sed -i s/localhost/jaeger/ ./packages/jaeger-ui/src/setupProxy.js
RUN yarn build

ENTRYPOINT [ "/usr/local/bin/yarn" ]

CMD [ "start" ]
