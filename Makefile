VOLTHA_TOOLS_VERSION ?= 2.3.1

PROTO_FILES := $(sort $(wildcard api/**/*.proto))

setup_tools:
	@echo "Downloading dependencies..."
	@go mod download github.com/grpc-ecosystem/grpc-gateway
	@echo "Dependencies downloaded OK"

#--go_out=./ --go_opt=paths=source_relative
protos: setup_tools
	protoc -I . \
		-I api \
		--go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		--openapiv2_out ./api/v1/gen/openapi \
		--openapiv2_opt logtostderr=true \
		--grpc-gateway_out ./api/v1/gen \
		--grpc-gateway_opt logtostderr=true \
		--grpc-gateway_opt paths=source_relative \
		$(PROTO_FILES)

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