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

## Setup

In order to play around with this code you need to deploy `aether-roc-umbrella` and
make sure you forward: 
- the `onos-config` gNMI server on port `5150` (`kubectl port-forward svc/onos-config 5150`)
- the `onos-topo` gRPC server on port `5151` (`kubectl port-forward svc/onos-topo 5151:5150`)


## Commands

The swagger UI for the REST APIs is available at: http://localhost:8080

Discover the gRPC services
```shell
grpcurl -plaintext localhost:50060 list
grpcurl -plaintext localhost:50060 describe roc.ApplicationApi
```

List all the applications via gRPC:
```shell
grpcurl -plaintext -d '{"enterpriseId": "acme"}' localhost:50060 roc.ApplicationApi/GetApplications
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

Via GraphQL (currently disabled)

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