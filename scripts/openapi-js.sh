#!/usr/bin/env bash
set -e

readonly service="$1"

rm -rf "web/src/repositories/clients/$service"
mkdir -p "web/src/repositories/clients/$service"

#docker run --rm --env "JAVA_OPTS=-Dlog.level=error" -v "${PWD}:/local" \
#  "openapitools/openapi-generator-cli:v4.3.0" generate \
#  -i "/local/api/openapi/$service.yml" \
#  -g javascript \
#  -o "/local/web/src/repositories/clients/$service"

docker run --rm --env "JAVA_OPTS=-Dlog.level=error" \
  -v "${PWD}:/local" "openapitools/openapi-generator-cli:v6.6.0" generate \
  -i "/local/api/openapi/$service.yml" \
  -g typescript-axios \
  -o "/local/web/src/repositories/clients/$service"
rm -rf "web/src/repositories/clients/$service/{.openapi-generator,.gitignore,.npmignore,.openapi-generator-ignore,git_push.sh}"