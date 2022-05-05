// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

//go:build tools
// +build tools

package tools

import (
	_ "github.com/99designs/gqlgen"
	_ "github.com/99designs/gqlgen/graphql/introspection"
	_ "github.com/danielvladco/go-proto-gql/protoc-gen-gogql"
	_ "github.com/danielvladco/go-proto-gql/protoc-gen-gql"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/google/gnostic/cmd/protoc-gen-openapi"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
