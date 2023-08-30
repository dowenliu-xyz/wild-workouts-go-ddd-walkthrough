include .env
export

# oapi-codegen 可以通过 go-generate 搞，不用写在 makefile 里
# protoc 可以换成 buf ，同时也支持在 go-generate 里搞，也不用写在 makefile 里
# 但是 openapi_js 生成 js 的命令不适合在 go-generate 里搞，跨项目类型了。
# 另外在 makefile 里统一写，有统一一条指令执行的特点。反之讲到 go-generate 或 npm 指令里需要分别按需要调用，可能生成的源 IDL 定义版本不一致。

.PHONY: openapi
openapi: openapi_http #openapi_js

.PHONY: openapi_http
# 把工具安装也写在依赖上。如果工具有版本的要求，也方便把版本固定下来。
# 这里依赖的 oapi-codegen 工具还没有正式版本，所以用 latest。有时有些依赖，更激进的，需要安装最新的不稳定的 master 版本。
openapi_http: tools.require.oapi-codegen
	mkdir -p internal/trainings/ports
	oapi-codegen -generate types -o internal/trainings/ports/openapi_types.gen.go -package ports api/openapi/trainings.yml
	oapi-codegen -generate chi-server -o internal/trainings/ports/openapi_api.gen.go -package ports api/openapi/trainings.yml
	mkdir -p internal/common/client/trainings
	oapi-codegen -generate types -o internal/common/client/trainings/openapi_types.gen.go -package trainings api/openapi/trainings.yml
	oapi-codegen -generate client -o internal/common/client/trainings/openapi_client_gen.go -package trainings api/openapi/trainings.yml

	mkdir -p internal/trainer/ports
	oapi-codegen -generate types -o internal/trainer/ports/openapi_types.gen.go -package ports api/openapi/trainer.yml
	oapi-codegen -generate chi-server -o internal/trainer/ports/openapi_api.gen.go -package ports api/openapi/trainer.yml
	mkdir -p internal/common/client/trainer
	oapi-codegen -generate types -o internal/common/client/trainer/openapi_types.gen.go -package trainer api/openapi/trainer.yml
	oapi-codegen -generate client -o internal/common/client/trainer/openapi_client_gen.go -package trainer api/openapi/trainer.yml

	mkdir -p internal/users
	oapi-codegen -generate types -o internal/users/openapi_types.gen.go -package main api/openapi/users.yml
	oapi-codegen -generate chi-server -o internal/users/openapi_api.gen.go -package main api/openapi/users.yml
	mkdir -p internal/common/client/users
	oapi-codegen -generate types -o internal/common/client/users/openapi_types.gen.go -package users api/openapi/users.yml
	oapi-codegen -generate client -o internal/common/client/users/openapi_client_gen.go -package users api/openapi/users.yml

#.PHONY: openapi_js
#openapi_js:
#	rm -rf web/src/repositories/clients
#	mkdir -p web/src/repositories/clients/trainings
#	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v6.6.0 generate \
#		-i /local/api/openapi/trainings.yml \
#		-g typescript-axios \
#		-o /local/web/src/repositories/clients/trainings
#	rm -rf web/src/repositories/clients/trainings/{.openapi-generator,.gitignore,.npmignore,.openapi-generator-ignore,git_push.sh}
#
#	mkdir -p web/src/repositories/clients/trainer
#	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v6.6.0 generate \
#		-i /local/api/openapi/trainer.yml \
#		-g typescript-axios \
#		-o /local/web/src/repositories/clients/trainer
#	rm -rf web/src/repositories/clients/trainer/{.openapi-generator,.gitignore,.npmignore,.openapi-generator-ignore,git_push.sh}
#
#	mkdir -p web/src/repositories/clients/users
#	docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli:v6.6.0 generate \
#		-i /local/api/openapi/users.yml \
#		-g typescript-axios \
#		-o /local/web/src/repositories/clients/users
#	rm -rf web/src/repositories/clients/users/{.openapi-generator,.gitignore,.npmignore,.openapi-generator-ignore,git_push.sh}

.PHONY: proto
proto:
	mkdir -p internal/common/genproto/trainer
	protoc --proto_path=api/protobuf \
		--go_out=internal/common/genproto/trainer --go_opt=paths=source_relative \
		--go-grpc_out=internal/common/genproto/trainer --go-grpc_opt=paths=source_relative \
		trainer.proto
	mkdir -p internal/common/genproto/users
	protoc --proto_path=api/protobuf \
		--go_out=internal/common/genproto/users --go_opt=paths=source_relative \
		--go-grpc_out=internal/common/genproto/users --go-grpc_opt=paths=source_relative \
		users.proto

.PHONY: lint
lint:
	@./scripts/lint.sh trainer
	@./scripts/lint.sh trainings
	@./scripts/lint.sh users

.PHONY: tools.require.oapi-codegen
tools.require.oapi-codegen:
	go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest

.PHONY: fmt
fmt:
	goimports -l -w internal/

.PHONY: mycli
mycli:
	mycli -u ${MYSQL_USER} -p ${MYSQL_PASSWORD} ${MYSQL_DATABASE}

.PHONY: test
test:
	@chmod +x ./scripts/test.sh
	@./scripts/test.sh common .e2e.env
	@./scripts/test.sh trainer .test.env
	@./scripts/test.sh trainings .test.env
	@./scripts/test.sh users .test.env
