//go:build tools
// +build tools

package tools

import (
	_ "github.com/danielvladco/go-proto-gql/protoc-gen-gogql"
	_ "github.com/danielvladco/go-proto-gql/protoc-gen-gql"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
