# syntax=docker/dockerfile:experimental

FROM golang:1.22-bookworm AS build-env

ARG TARGETARCH

# Only install Delve if architecture is NOT s390x
RUN if [ "$TARGETARCH" != "s390x" ]; then \
      go install github.com/go-delve/delve/cmd/dlv@v1.21.0 ; \
    else \
      echo "Skipping Delve installation for unsupported arch: $TARGETARCH" ; \
    fi

FROM debian:bookworm-slim

RUN apt update && \
    apt install procps -y && \
    rm -rf /var/lib/apt/lists/*

RUN if [ -f /go/bin/dlv ]; then cp /go/bin/dlv /dlv; else echo "Skipping dlv copy"; fi

CMD ["bash", "-c", "/dlv attach $(ps axf | grep $CMD_NAME | grep -v grep | awk '{print $1}') --listen=:8000 --headless=true --api-version=2 --accept-multiclient --continue"]
