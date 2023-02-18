include .env
export

.PHONY: aqua
aqua: # export PATH="${AQUA_ROOT_DIR:-${XDG_DATA_HOME:-$HOME/.local/share}/aquaproj-aqua}/bin:$PATH"
	@go run github.com/aquaproj/aqua-installer@latest

.PHONY: tool
tool:
	@aqua i
	@(cd api && npm install)

.PHONY: doc
doc:
	@rm -rf doc
	@tbls doc ${DATABASE_URL} doc/database
	@mkdir -p doc/proto
	@protoc --doc_out=./doc/proto --doc_opt=markdown,README.md \
		proto/article/**/*.proto \
		proto/auth/**/*.proto \
		proto/health/**/*.proto
	@npx widdershins --omitHeader --code true ./api/openapi.yaml doc/openapi.md

.PHONY: lint
lint:
	@golangci-lint run --fix

.PHONY: modules
modules:
	@go list -u -m all

.PHONY: renovate
renovate:
	@go get -u -t ./...

.PHONY: compile
compile:
	@go build -v ./... && go clean

.PHONY: test
test:
	@go test ./internal/...

.PHONY: e2e
e2e:
	@go test ./e2e/... -tags e2e -count=1

.PHONY: gen
gen:
	@go generate ./...
	@oapi-codegen -generate types -package openapi ./api/openapi.yaml > ./pkg/openapi/types.gen.go
	@oapi-codegen -generate chi-server -package openapi ./api/openapi.yaml > ./pkg/openapi/server.gen.go
	@oapi-codegen -generate client -package openapi ./api/openapi.yaml > ./pkg/openapi/client.gen.go
	@buf generate --template ./proto/buf.gen.yaml
	@go mod tidy

.PHONY: bufmt
bufmt:
	@buf format -w

.PHONY: buflint
buflint:
	@(cd proto && buf lint)

.PHONY: apilint
apilint:
	@(cd api && npx spectral lint openapi.yaml)

.PHONY: ymlint
ymlint:
	@yamlfmt -lint && actionlint

.PHONY: ymlfmt
ymlfmt:
	@yamlfmt

.PHONY: dup
dup:
	@docker compose --project-name ${APP_NAME} --file ./.docker/docker-compose.yaml up -d

.PHONY: ddown
ddown:
	@docker compose --project-name ${APP_NAME} down

.PHONY: psql
psql:
	@docker exec -it ${APP_NAME}-postgres psql -U postgres
