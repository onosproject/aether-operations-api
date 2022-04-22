# Roc API

## Setup

In order to play around with this code you need to deploy `aether-roc-umbrella` and
make sure you forward  the `onos-config` gNMI server on port `5150`.

## Commands

List all the applications via gRPC:
```shell
grpcurl -plaintext localhost:50060 roc.RocApi/GetApplications
```

List all the applciations via REST:
```shell
curl http://localhost:8181/api/v1/applications
```