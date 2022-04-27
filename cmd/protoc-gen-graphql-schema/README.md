Install with:
```shell
go install ./cmd/protoc-gen-graphql-schema 
```

Generate protos with:
```shell
protoc -I .\
    -I api \
    -I vendor/github.com/grpc-ecosystem/grpc-gateway/v2/ \
    --graphql-schema_out=./api/v1 --graphql-schema_opt=paths=source_relative\
    api/v1/*.proto
```