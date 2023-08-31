#!/usr/bin/env bash

if [ -x "$(go env GOPATH)/bin/golangci-lint" ]; then
    exit 0
else
  set -ex
  curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.54.1
fi
