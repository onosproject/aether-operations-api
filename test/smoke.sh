#!/usr/bin/env bash

echo "Testing gRPC server"

grpcurl -plaintext localhost:50060 list
grpcurl -plaintext localhost:50060 describe application.ApplicationService
grpcurl -plaintext -d '{"enterpriseId": "acme"}' localhost:50060 application.ApplicationService/GetApplications

echo "--- done testing gRPC"

echo "Testing REST server"

curl -X 'GET' \
  'http://localhost:8080/api/v1/enterpise/acme/applications' \
  -H 'accept: application/json' | jq .
echo " "
curl -X 'GET' \
  'http://localhost:8080/api/v1/applications?enterpriseId=acme' \
  -H 'accept: application/json' | jq .

echo -e "\n--- done testing REST"

echo "Testing GraphQL server"

curl 'http://localhost:8081/application-query' \
    -H 'accept: application/json, multipart/mixed' \
    -H 'content-type: application/json' \
    --data-raw '{"query":"query {\n  applicationServiceGetApplications(in: {enterpriseId: \"acme\"} ) {\n    applications {\n      iD\n      endpoint {\n        iD\n       displayName\n}\n    }\n  }\n}"}' \
    | jq .

curl 'http://localhost:8081/enterprise-query' \
  -H 'accept: application/json, multipart/mixed' \
  -H 'content-type: application/json' \
  --data-raw '{"query":"query {\n  enterpriseServiceGetEnterprises {\n    enterprises {\n      iD\n    }\n  }\n}","variables":null}' \
  | jq .

echo -e "\n--- done testing GraphQL"