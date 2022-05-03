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

Requires `protoc` (for now) and `buf` to be installed.

## Local Setup

In order to play around with this code you need to deploy `aether-roc-umbrella` and
make sure you forward: 
- the `onos-config` gNMI server on port `5150` (`kubectl port-forward svc/onos-config 5150`)
- the `onos-topo` gRPC server on port `5151` (`kubectl port-forward svc/onos-topo 5151:5150`)

## Deployment

If you want to deploy `scaling-umbrella` in you cluster you can use the provided helm chart.
Since the image is not yet available on DockerHub you need to make it available to your cluster.

You can build and load the image to a `Kind` cluster with `make kind`.
If you are using a different setup you can build the docker image with `make docker-build` and load it in you cluster as appropriate.

```shell
 helm upgrade --install scaling-umbrella ./deployments/scaling-umbrella 
```

> NOTE the default `values` for this chart assume that it is deployed in the same `namespace` as `aether-roc-umbrella`
> 
> If you deploy in a different namespace you can customize the `onos-config` and `onos-topo` endpoints with:
> ```shell
> helm upgrade --install scaling-umbrella ./deployments/scaling-umbrella \
>   --set scalingUmbrella.dataSources.onosConfig=onos-config.<namespace>.svc:5150 \
>   --set scalingUmbrella.dataSources.onosTopo=onos-topo.<namespace>.svc:5150
> ```

_For more information on the supported customization refer to the [values.yaml](deployments/scaling-umbrella/values.yaml) file._

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
