# syntax=docker/dockerfile:experimental

FROM debian:bookworm-slim

ENV DEBIAN_FRONTEND noninteractive

ARG HTTPS_PROXY
ARG HTTP_PROXY

ENV http_proxy $HTTP_PROXY
ENV https_proxy $HTTPS_PROXY

RUN apt update && \
    apt install build-essential curl git pkg-config libfuse-dev fuse -y && \
    rm -rf /var/lib/apt/lists/*

RUN curl -fsSL https://deb.nodesource.com/setup_18.x | bash - && \
    apt install -y nodejs && \
    rm -rf /var/lib/apt/lists/*
RUN npm install pnpm@9 -g

ARG TARGET_PLATFORM=amd64

RUN curl -vvv https://dl.google.com/go/go1.22.10.linux-$TARGET_PLATFORM.tar.gz | tar -xz -C /usr/local
ENV PATH "/usr/local/go/bin:${PATH}"
ENV GOCACHE /tmp/go-build
ENV GOPATH /tmp/go

LABEL org.opencontainers.image.source https://github.com/chaos-mesh/chaos-mesh
