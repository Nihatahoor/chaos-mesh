# syntax=docker/dockerfile:experimental


FROM debian:bookworm-slim

ENV DEBIAN_FRONTEND=noninteractive

ARG HTTPS_PROXY
ARG HTTP_PROXY

ENV http_proxy=$HTTP_PROXY
ENV https_proxy=$HTTPS_PROXY

RUN apt update && \
    apt install -y build-essential curl git pkg-config libfuse-dev fuse && \
    rm -rf /var/lib/apt/lists/*

RUN npm_install() { \
    ARCH=$(dpkg --print-architecture); \
    if [ "$ARCH" = "s390x" ]; then \
      echo "Installing Node.js manually for s390x"; \
      curl -fsSL https://nodejs.org/dist/v18.18.1/node-v18.18.1-linux-s390x.tar.xz | tar -xJ -C /usr/local --strip-components=1; \
    else \
      apt update && apt install -y nodejs npm; \
    fi; \
}; npm_install

RUN npm install -g pnpm@9

ARG TARGET_PLATFORM=s390x

RUN curl -vvv https://dl.google.com/go/go1.22.10.linux-$TARGET_PLATFORM.tar.gz | tar -xz -C /usr/local

ENV PATH="/usr/local/go/bin:${PATH}"
ENV GOCACHE=/tmp/go-build
ENV GOPATH=/tmp/go

LABEL org.opencontainers.image.source https://github.com/chaos-mesh/chaos-mesh
