include .env
export

# 在 makefile 里统一写，有统一使用同一版本生成各端代码的特点。防止各模块各自生成时，引用 IDL 版本不一致问题。

.PHONY: openapi
openapi: openapi_http #openapi_js

.PHONY: openapi_http
# 把工具安装也写在依赖上。如果工具有版本的要求，也方便把版本固定下来。
# 这里依赖的 oapi-codegen 工具还没有正式版本，所以用 latest。有时有些依赖，更激进的，需要安装最新的不稳定的 master 版本。
openapi_http: tools.require.oapi-codegen
	@chmod +x ./scripts/openapi-http.sh
	@./scripts/openapi-http.sh trainer internal/trainer/ports ports
	@./scripts/openapi-http.sh trainings internal/trainings/ports ports
	@./scripts/openapi-http.sh users internal/users main

#.PHONY: openapi_js
#openapi_js:
#	@chmod +x ./scripts/openapi-js.sh
#	@./scripts/openapi-js.sh trainer
#	@./scripts/openapi-js.sh trainings
#	@./scripts/openapi-js.sh users

.PHONY: proto
proto:
	@chmod +x ./scripts/proto.sh
	@./scripts/proto.sh trainer
	@./scripts/proto.sh users

.PHONY: lint
lint: tools.require.go-cleanarch tools.require.golangci-lint
	@go-cleanarch
	@./scripts/lint.sh common
	@./scripts/lint.sh trainer
	@./scripts/lint.sh trainings
	@./scripts/lint.sh users

.PHONY: tools.require.%
tools.require.%:
	@chmod +x ./scripts/tools/$*.sh && ./scripts/tools/$*.sh


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
