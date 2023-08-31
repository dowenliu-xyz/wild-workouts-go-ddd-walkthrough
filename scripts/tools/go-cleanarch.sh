#!/usr/bin/env bash

if [ -x "$(go env GOPATH)/bin/go-cleanarch" ]; then
    exit 0
else
  set -ex
  go install github.com/roblaszczak/go-cleanarch@latest
fi
