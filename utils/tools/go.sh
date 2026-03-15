#!/bin/bash


export GO_VERSION="1.26.1"
export GOLG_DNLOAD_LINK="https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz"

curl -LO ${GOLG_DNLOAD_LINK}

if [ -f /usr/local/go ]; then
  sudo rm -rf /usr/local/go
fi;

sudo tar -C /usr/local -xzf go${GO_VERSION}.linux-amd64.tar.gz
sudo ln -s /usr/local/go/bin/go /usr/bin/go

go version