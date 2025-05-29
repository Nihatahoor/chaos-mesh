FROM curlimages/curl:7.88.1 as binaryimage

USER root

ARG HTTPS_PROXY
ARG HTTP_PROXY

ENV http_proxy $HTTP_PROXY
ENV https_proxy $HTTPS_PROXY

ARG TARGET_PLATFORM=s390x

RUN mkdir -p /tmp/bin && \
    curl -L https://mirrors.chaos-mesh.org/byteman-chaos-mesh-download-v4.0.24-0.12.tar.gz -o /usr/local/byteman.tar.gz && \
    tar xvf /usr/local/byteman.tar.gz -C /usr/local && \
    mv /usr/local/byteman-chaos-mesh-download-v4.0.24-0.12 /tmp/byteman && \
    rm /usr/local/byteman.tar.gz

# 🛑 Optional: only include toda if supported (amd64 only)
RUN if [ "$TARGET_PLATFORM" = "amd64" ]; then \
    curl -L https://github.com/chaos-mesh/toda/releases/download/v0.2.4/toda-linux-amd64.tar.gz | tar xz -C /tmp/bin; \
    fi

# ✅ Multi-arch case block (with s390x support)
RUN case "$TARGET_PLATFORM" in \
    amd64) export NSEXEC_ARCH="x86_64" ;; \
    arm64) export NSEXEC_ARCH="aarch64" ;; \
    s390x) export NSEXEC_ARCH="s390x" ;; \
    *) echo >&2 "Unsupported arch '$TARGET_PLATFORM'"; exit 1 ;; \
    esac && \
    curl -L https://github.com/chaos-mesh/nsexec/releases/download/v0.1.6/nsexec-$NSEXEC_ARCH-unknown-linux-gnu.tar.gz | tar xz -C /tmp/bin || true && \
    curl -L https://github.com/chaos-mesh/chaos-tproxy/releases/download/v0.5.3/tproxy-$NSEXEC_ARCH.tar.gz | tar xz -C /tmp/bin || true && \
    curl -L https://github.com/chaos-mesh/memStress/releases/download/v0.3/memStress_v0.3-$NSEXEC_ARCH-linux-gnu.tar.gz | tar xz -C /tmp/bin || true


# ---
FROM debian:bookworm-slim

ARG HTTPS_PROXY
ARG HTTP_PROXY

ENV http_proxy $HTTP_PROXY
ENV https_proxy $HTTPS_PROXY

RUN mkdir -p /usr/share/man/man1 /usr/share/man/man2

RUN apt-get update && \
    apt-get install -y tzdata iptables ebtables net-tools ipset stress-ng iproute2 fuse util-linux procps default-jre && \
    rm -rf /var/lib/apt/lists/*



RUN update-alternatives --set iptables /usr/sbin/iptables-legacy && \
    update-alternatives --set ebtables /usr/sbin/ebtables-legacy

# ✅ This might differ on s390x — double check with `readlink -f $(which java)`
ENV RUST_BACKTRACE=1
ENV BYTEMAN_HOME=/usr/local/byteman
ENV PATH=$PATH:/usr/local/byteman/bin:/usr/local/bin
ENV JAVA_HOME=/usr/lib/jvm/java-11-openjdk-s390x

# ✅ Your s390x compiled binaries
COPY bin/chaos-daemon /usr/local/bin/chaos-daemon
COPY bin/pause /usr/local/bin/pause
#COPY bin/cdh /usr/local/bin/cdh

COPY --from=binaryimage /tmp/byteman /usr/local/byteman
#COPY --from=binaryimage /tmp/bin/* /usr/local/bin/
COPY --from=binaryimage /tmp/bin/ /usr/local/bin/

RUN mv /usr/local/bin/libnsenter.so /usr/local/lib/libnsenter.so || true

