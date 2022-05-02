# Roc API

## NOTES

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

* Writing the graphql schema from hand might actually be preferred for now. Esp. if GUI development has yet to be defined

## Required tools

Requires `protoc` to be installed.

## Setup

In order to play around with this code you need to deploy `aether-roc-umbrella` and
make sure you forward: 
- the `onos-config` gNMI server on port `5150` (`kubectl port-forward svc/onos-config 5150`)
- the `onos-topo` gRPC server on port `5151` (`kubectl port-forward svc/onos-topo 5151:5150`)


## Usage

The swagger UI for the REST APIs is available at: http://localhost:8080
The graphQL playground is available at:
- http://localhost:8081/enterprise-playground
- http://localhost:8081/application-playground

For more usage examples see `test/smoke.sh` 

## Buf

Requires buf installed eg. `brew install buf`

* `make buf` - mod update, build, generate api
* other stuff
  * `buf build --exclude-source-info -o -#format=json | jq '.file[] | .package'`
  * `buf ls-files`
  * `buf lint`
