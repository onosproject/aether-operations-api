# Roc API

## NOTEs

A strawman project to generate:
- REST APIs
- Swagger Docs
- GraphQL server
starting from Protobuf.

### Caveats and TODOs

- The tools used to generate GraphQL doesn't seem to be very mature/supported.
  - We might consider to investigate https://github.com/google/rejoiner or https://github.com/bi-foundation/protobuf-graphql-extension
  - Ideally we want something that can generate the schema and we can write our own resolvers.
- Needs some work to share the REST server between REST API and GraphQL

## Required tools

Requires `protoc` to be installed.

Addictionally requires some `go` packages to be installed:

```shell
go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
    github.com/ysugimoto/grpc-graphql-gateway/protoc-gen-graphql
```

## Setup

In order to play around with this code you need to deploy `aether-roc-umbrella` and
make sure you forward  the `onos-config` gNMI server on port `5150`.

## Commands

The swagger UI for the REST APIs is available at: http://localhost:8080

List all the applications via gRPC:
```shell
grpcurl -plaintext -d '{"enterpriseId": "acme"}' localhost:50060 roc.RocApi/GetApplications
# {
#   "applications": [
#     {
#       "ID": "acme-dataacquisition",
#       "description": "Data Acquisition",
#       "endpoint": [
#         {
#           "ID": "da",
#           "DisplayName": "data acquisition endpoint",
#           "Mbr": {
#             "uplink": "2000000",
#             "downlink": "1000000"
#           },
#           "PortStart": 7585,
#           "PortEnd": 7588,
#           "Protocol": "TCP"
#         }
#       ]
#     }
#   ]
# }
```

List Applications for a specific Enterprise via REST:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/enterpise/acme/applications' \
  -H 'accept: application/json'
# {"applications":[{"ID":"acme-dataacquisition", "description":"Data Acquisition", "endpoint":[{"ID":"da", "Description":"", "DisplayName":"data acquisition endpoint", "Mbr":{"uplink":"2000000", "downlink":"1000000"}, "PortStart":7585, "PortEnd":7588, "Protocol":"TCP"}]}]}     
```
or
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/applications?enterpriseId=acme' \
  -H 'accept: application/json'
# {"applications":[{"ID":"acme-dataacquisition", "description":"Data Acquisition", "endpoint":[{"ID":"da", "Description":"", "DisplayName":"data acquisition endpoint", "Mbr":{"uplink":"2000000", "downlink":"1000000"}, "PortStart":7585, "PortEnd":7588, "Protocol":"TCP"}]}]} 
```

Via GraphQL

```shell
curl -g "http://localhost:8081/graphql" -d '
{
  getApplications(enterpriseId: "acme") {
    applications {
      ID,
      description
    }
  }
}'
# {"data":{"getApplications":{"applications":[{"ID":"acme-dataacquisition","description":"Data Acquisition"}]}}}
```