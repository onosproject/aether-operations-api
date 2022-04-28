# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

PROTO_FILES := $(sort $(wildcard api/**/*.proto))

help: # @HELP Print the command options
	@echo
	@echo "\033[0;31m    ROC API (scaling-umbrella) \033[0m"
	@echo
	@echo Aether ROC APIs
	@echo
	@grep -E '^.*: .* *# *@HELP' $(MAKEFILE_LIST) \
    | sort \
    | awk ' \
        BEGIN {FS = ": .* *# *@HELP"}; \
        {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}; \
    '

setup_tools: mod-update
	@echo "Downloading dependencies..."
	go install \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
        github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc \
        github.com/danielvladco/go-proto-gql/protoc-gen-gql \
        github.com/danielvladco/go-proto-gql/protoc-gen-gogql
	@echo "Dependencies downloaded OK"

protos: setup_tools # @HELP Generates Go Models, gRPC Interface, REST Gateway and Swagger APIs
	protoc -I . \
		-I api \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
		-I vendor/github.com/danielvladco/go-proto-gql/ \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--openapiv2_out ./api/swagger/dist \
		--openapiv2_opt logtostderr=true,generate_unbound_methods=true \
		--openapiv2_opt openapi_naming_strategy=simple \
		--openapiv2_opt allow_merge=true,merge_file_name=roc,output_format=yaml \
		--gql_out=paths=source_relative:. \
		--gogql_out=paths=source_relative:. \
		$(PROTO_FILES)

graphql:
	# FIXME looks like gqlgen ignores the config file name and always reads gqlgen.yaml
	cp internal/servers/graphql/config/gqlgen.apps.yaml gqlgen.yaml
	go run github.com/99designs/gqlgen --config gqlgen.apps.yaml --verbose generate
	cp internal/servers/graphql/config/gqlgen.ent.yaml gqlgen.yaml
	go run github.com/99designs/gqlgen --config gqlgen.ent.yaml --verbose generate
	rm gqlgen.yaml

.PHONY: build
build: protos graphql build-go # @HELP Build the protos, graphql gateway and go executable

build-go: # @HELP Build the go executable
	@go build -mod vendor \
	  -ldflags "-w -X main.buildTime=$(date +%Y/%m/%d-%H:%M:%S) \
		-X main.commitHash=$(git log --pretty=format:%H -n 1) \
		-X main.gitStatus=${GIT_STATUS} \
		-X main.version=${VERSION}" \
	  ./cmd/roc-api

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
	golangci-lint run

mod-update: # @HELP Download the dependencies to the vendor folder
	go mod tidy
	go mod vendor