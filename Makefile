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
        github.com/danielvladco/go-proto-gql/protoc-gen-gogql \
        github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql
	@echo "Dependencies downloaded OK"

protos: setup_tools # @HELP Generates Go Models, gRPC Interface, REST Gateway and Swagger APIs
	protoc -I . \
		-I api \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out . \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		--openapiv2_out ./api/swagger/dist \
		--openapiv2_opt logtostderr=true,generate_unbound_methods=true \
		--openapiv2_opt openapi_naming_strategy=simple \
		--openapiv2_opt allow_merge=true,merge_file_name=roc,output_format=yaml \
		$(PROTO_FILES)

schema: # @HELP [Experimental] Generates GraphQL schema using github.com/danielvladco/go-proto-gql
	protoc -I .\
		-I api \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
		--gql_out=paths=source_relative:. $(PROTO_FILES)

schema-local: # @HELP [Super-Experimental] Generates GraphQL schema using a custom plugin (defined in cmd/protoc-gen-graphql-schema)
	go install ./cmd/protoc-gen-graphql-schema && protoc -I .\
		-I api \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
		--graphql-schema_out=./api/v1 --graphql-schema_opt=paths=source_relative\
    	$(PROTO_FILES)


graphql: # @HELP [Experimental] Generates a GraphQL Gateway using github.com/ysugimoto/grpc-graphql-gateway
	# generate the types
	protoc -I .\
		-I api \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
		--graphql_opt=paths=source_relative \
		--graphql_out=./api/v1 \
		./api/v1/application.proto
	mv api/v1/v1.graphql.go api/v1/application.graphql.go

	protoc -I .\
		-I api \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
		--graphql_opt=paths=source_relative \
		--graphql_out=./api/v1 \
		./api/v1/endpoint.proto
	mv api/v1/v1.graphql.go api/v1/endpoint.graphql.go

	protoc -I .\
		-I api \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
		--graphql_opt=paths=source_relative \
		--graphql_out=./api/v1 \
		./api/v1/enterprise.proto
	mv api/v1/v1.graphql.go api/v1/enterprise.graphql.go

	protoc -I .\
		-I api \
		-I vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
		--graphql_opt=paths=source_relative \
		--graphql_out=./api/v1 \
		./api/v1/roc.proto
	mv api/v1/v1.graphql.go api/v1/roc.graphql.go



.PHONY: build
build: protos # @HELP Build the go executable
	@go build -mod vendor \
	  -ldflags "-w -X main.buildTime=$(date +%Y/%m/%d-%H:%M:%S) \
		-X main.commitHash=$(git log --pretty=format:%H -n 1) \
		-X main.gitStatus=${GIT_STATUS} \
		-X main.version=${VERSION}" \
	  ./cmd/roc-api

mod-update: # @HELP Download the dependencies to the vendor folder
	go mod tidy
	go mod vendor