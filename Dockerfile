# Build Go binary without cgo dependencies
FROM golang:1.13.4 as builder
WORKDIR /go/src/github.com/gocardless/theatre
COPY . /go/src/github.com/gocardless/theatre
RUN make VERSION=$(cat VERSION) build
# Clone our fork of envconsul and build it
RUN git clone https://github.com/gocardless/envconsul.git \
      && make -C envconsul linux/amd64 \
      && mv envconsul/pkg/linux_amd64/envconsul bin

# Use ubuntu as our base package to enable generic system tools
FROM ubuntu:bionic-20191202

# Without these certificates we'll fail to validate TLS connections to Google's
# services.
RUN set -x \
      && apt-get update -y \
      && apt-get install -y --no-install-recommends \
                            ca-certificates \
      && apt-get clean -y \
      && rm -rf /var/lib/apt/lists/*

WORKDIR /
COPY --from=builder /go/src/github.com/gocardless/theatre/bin/* /usr/local/bin/
ENTRYPOINT ["/bin/bash"]
