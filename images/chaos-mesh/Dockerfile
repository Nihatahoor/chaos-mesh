FROM --platform=linux/s390x alpine:3.21

ARG HTTPS_PROXY
ARG HTTP_PROXY

RUN apk add --no-cache tzdata

COPY images/chaos-mesh/bin/chaos-controller-manager /usr/local/bin/chaos-controller-manager

ENTRYPOINT ["/usr/local/bin/chaos-controller-manager"]







