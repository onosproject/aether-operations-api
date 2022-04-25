VOLTHA_TOOLS_VERSION ?= 2.3.1

PROTO_FILES := $(sort $(wildcard api/**/*.proto))

setup_tools:
	@echo "Downloading dependencies..."
	@go mod download github.com/grpc-ecosystem/grpc-gateway
	@echo "Dependencies downloaded OK"

protos: setup_tools
	## TODO decouple APIS
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

graphql:
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
build: protos
	@go build -mod vendor \
	  -ldflags "-w -X main.buildTime=$(date +%Y/%m/%d-%H:%M:%S) \
		-X main.commitHash=$(git log --pretty=format:%H -n 1) \
		-X main.gitStatus=${GIT_STATUS} \
		-X main.version=${VERSION}" \
	  ./cmd/roc-api

mod-update: # @HELP Download the dependencies to the vendor folder
	go mod tidy
	go mod vendor