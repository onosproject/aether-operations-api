# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

SHELL = bash -e -o pipefail
VERSION     ?= $(shell cat ./VERSION)
PROTO_FILES := $(sort $(wildcard api/**/*.proto))

DOCKER_TAG  			?= ${VERSION}
DOCKER_REPOSITORY  		?= onosproject/
DOCKER_REGISTRY 		?=
## Docker labels. Only set ref and commit date if committed
DOCKER_LABEL_VCS_URL       ?= $(shell git remote get-url $(shell git remote))
DOCKER_LABEL_VCS_REF       = $(shell git rev-parse HEAD)
DOCKER_LABEL_BUILD_DATE    ?= $(shell date -u "+%Y-%m-%dT%H:%M:%SZ")
DOCKER_LABEL_COMMIT_DATE   = $(shell git show -s --format=%cd --date=iso-strict HEAD)
DOCKER_LABEL_VCS_DIRTY     = false
ifneq ($(shell git ls-files --others --modified --exclude-standard 2>/dev/null | wc -l | sed -e 's/ //g'),0)
    DOCKER_LABEL_VCS_DIRTY = true
endif
DOCKER_BUILD_ARGS ?= \
	${DOCKER_ARGS} \
	--build-arg org_label_schema_version="${VERSION}" \
	--build-arg org_label_schema_vcs_url="${DOCKER_LABEL_VCS_URL}" \
	--build-arg org_label_schema_vcs_ref="${DOCKER_LABEL_VCS_REF}" \
	--build-arg org_label_schema_build_date="${DOCKER_LABEL_BUILD_DATE}" \
	--build-arg org_opencord_vcs_commit_date="${DOCKER_LABEL_COMMIT_DATE}" \
	--build-arg org_opencord_vcs_dirty="${DOCKER_LABEL_VCS_DIRTY}"
KIND_CLUSTER_NAME 		?= kind

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

buf: # @HELP Generates Go Models, gRPC Interface, REST Gateway and Swagger APIs
	buf mod update api
	buf build
	buf generate api

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

# TODO: consider just crafting a single graphql schema and break it down only if needed
# or figure out way to stitch them together... discuss tradeoffs
# * doesn't need 1:1 mapping of proto models and maybe the decoupling is good
# * simpler protos
# * may allow for more idiomatic implementations
gql:
	go run github.com/99designs/gqlgen --verbose generate

graphql:
	# FIXME looks like gqlgen ignores the config file name and always reads gqlgen.yaml
	cp internal/servers/graphql/config/gqlgen.apps.yaml gqlgen.yaml
	go run github.com/99designs/gqlgen --config gqlgen.apps.yaml --verbose generate
	cp internal/servers/graphql/config/gqlgen.ent.yaml gqlgen.yaml
	go run github.com/99designs/gqlgen --config gqlgen.ent.yaml --verbose generate
	rm gqlgen.yaml

.PHONY: build
build: protos graph build-go # @HELP Build the protos, graphql gateway and go executable

build-go: # @HELP Build the go executable
	@go build -mod vendor \
	  -ldflags "-w -X github.com/onosproject/scaling-umbrella/internal/config.buildTime=${DOCKER_LABEL_BUILD_DATE} \
		-X github.com/onosproject/scaling-umbrella/internal/config.commitHash=${DOCKER_LABEL_VCS_REF} \
		-X github.com/onosproject/scaling-umbrella/internal/config.vcsDirty=${DOCKER_LABEL_VCS_DIRTY} \
		-X github.com/onosproject/scaling-umbrella/internal/config.version=${VERSION}" \
	  ./cmd/roc-api

lint-go:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.45.2
	golangci-lint run

mod-update: # @HELP Download the dependencies to the vendor folder
	go mod tidy
	go mod vendor

docker-build: # @HELP Build the BBSim docker container (contains BBSimCtl too)
	docker build \
	  ${DOCKER_BUILD_ARGS} \
	  -t ${DOCKER_REGISTRY}${DOCKER_REPOSITORY}scaling-umbrella:${DOCKER_TAG} \
	  -f build/Dockerfile .

docker-push: # @HELP Push the docker container to a registry
	docker push ${DOCKER_REGISTRY}${DOCKER_REPOSITORY}scaling-umbrella:${DOCKER_TAG}

kind-only: # @HELP Load the docker container into a kind cluster (cluster name can be customized with KIND_CLUSTER_NAME)
	@if [ "`kind get clusters`" = '' ]; then echo "no kind cluster found" && exit 1; fi
	kind load docker-image --name ${KIND_CLUSTER_NAME} ${DOCKER_REPOSITORY}scaling-umbrella:${DOCKER_TAG}

kind: docker-build kind-only
