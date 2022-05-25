# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0

SHELL = bash -e -o pipefail
VERSION     ?= $(shell cat ./VERSION)

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
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc \
        github.com/danielvladco/go-proto-gql/protoc-gen-gql \
        github.com/danielvladco/go-proto-gql/protoc-gen-gogql \
        github.com/google/gnostic/cmd/protoc-gen-openapi
	@echo "Dependencies downloaded OK"

buf: # @HELP Generates Go Models, gRPC Interface, REST Gateway and Swagger APIs
	buf mod update api/v1
	buf build
	buf generate

gql:
	go run github.com/99designs/gqlgen --verbose generate --config internal/servers/graphql/gqlgen.yaml

clean-gen: # @HELP Removes all generated files
	rm -r ./gen/go || true
	rm -r ./gen/graph || true

PROTO_DIR = api/v1
lint-proto: $(PROTO_DIR)/*	# @HELP Runs a lint and backward compatibility on the protos
	@for file in $^ ; do \
		if [ -d "$${file}" ]; then \
			echo "Linting" $${file} ; \
			buf lint --path $${file}; \
		fi; \
	done
	buf breaking --against '.git#branch=main'

.PHONY: build
build: local-aether-models setup_tools buf gql build-go # @HELP Build the protos, graphql gateway and go executable

local-aether-models: ## Copies a local version of the aether-models dependency into the vendor directory
ifdef LOCAL_AETHER_MODELS
	rm -rf vendor/github.com/onosproject/aether-models
	mkdir -p vendor/github.com/onosproject/aether-models/models/aether-2.1.x/v2/api
	cp -r ${LOCAL_AETHER_MODELS}/models/aether-2.1.x/api/* vendor/github.com/onosproject/aether-models/models/aether-2.1.x/v2/api
endif

build-go: local-aether-models # @HELP Build the go executable
	@go build -mod vendor -v \
	  -ldflags "-w -X github.com/onosproject/scaling-umbrella/internal/config.buildTime=${DOCKER_LABEL_BUILD_DATE} \
		-X github.com/onosproject/scaling-umbrella/internal/config.commitHash=${DOCKER_LABEL_VCS_REF} \
		-X github.com/onosproject/scaling-umbrella/internal/config.vcsDirty=${DOCKER_LABEL_VCS_DIRTY} \
		-X github.com/onosproject/scaling-umbrella/internal/config.version=${VERSION}" \
	  ./cmd/roc-api

lint-go: # @HELP Lints the GO code
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
