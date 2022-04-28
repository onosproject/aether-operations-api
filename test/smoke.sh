#!/usr/bin/env bash

echo "Testing gRPC server"

grpcurl -plaintext localhost:50060 list
grpcurl -plaintext localhost:50060 describe application.ApplicationService
grpcurl -plaintext -d '{"enterpriseId": "acme"}' localhost:50060 application.ApplicationService/GetApplications

echo "--- done testing gRPC"

echo "Testing REST server"

curl -X 'GET' \
  'http://localhost:8080/api/v1/enterpise/acme/applications' \
  -H 'accept: application/json'
echo " "
curl -X 'GET' \
  'http://localhost:8080/api/v1/applications?enterpriseId=acme' \
  -H 'accept: application/json'

echo -e "\n--- done testing REST"

echo "Testing GraphQL server"

curl -X POST "http://localhost:8081/graphql" -d '
{
  getApplications(enterpriseId: "acme") {
    applications {
      ID,
      description
    }
  }
}'

echo "--- done testing GraphQL"