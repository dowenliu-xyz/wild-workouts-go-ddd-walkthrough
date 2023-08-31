#!/usr/bin/env bash

if [ -x "$(go env GOPATH)/bin/oapi-codegen" ]; then
    exit 0
else
  set -ex
  go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
fi
