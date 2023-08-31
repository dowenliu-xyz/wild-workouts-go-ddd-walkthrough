#!/usr/bin/env bash

export FIRESTORE_EMULATOR_HOST=fake:8787
export TRAINER_GRPC_ADDR=fake:3000
export USERS_GRPC_ADDR=fake:3001
export GCP_PROJECT=fake

rm -rf out
mkdir out

go run .

if ! type "plantuml" > /dev/null; then
  echo "Please install plantuml to generate PNG diagrams automatically"
  exit 1
fi

for f in out/*.plantuml; do plantuml "$f"; done
