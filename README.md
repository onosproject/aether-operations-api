# Roc API

## Setup

In order to play around with this code you need to deploy `aether-roc-umbrella` and
make sure you forward  the `onos-config` gNMI server on port `5150`.

## Commands

List all the applications via gRPC:
```shell
grpcurl -plaintext localhost:50060 roc.RocApi/GetApplications
```

List Applications for a specific Enterprise via REST:
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/enterpise/acme/applications' \
  -H 'accept: application/json'
```
or
```shell
curl -X 'GET' \
  'http://localhost:8080/api/v1/applications?enterpriseId=acme' \
  -H 'accept: application/json'
```