#!/usr/bin/env bash

echo "Testing gRPC server"

grpcurl -plaintext localhost:50060 list
grpcurl -plaintext localhost:50060 describe applications.v1.ApplicationService
grpcurl -plaintext localhost:50060 enterprises.v1.EnterpriseService/GetEnterprises
grpcurl -plaintext -d '{"enterpriseId": "acme"}' localhost:50060 applications.v1.ApplicationService/GetApplications
grpcurl -plaintext -d '{"enterpriseId": "acme"}' localhost:50060 sites.v1.SiteService/GetSites

echo "--- done testing gRPC"

echo "Testing REST server"

curl --fail -X 'GET' \
  'http://localhost:8080/api/v1/enterpise/acme/applications' \
  -H 'accept: application/json' | jq .
echo " "
curl --fail -X 'GET' \
  'http://localhost:8080/api/v1/applications?enterpriseId=acme' \
  -H 'accept: application/json' | jq .

curl --fail -X 'GET' \
  'http://localhost:8080/api/v1/sites?enterprise_id=acme' \
  -H 'accept: application/json' | jq .

curl --fail -X 'GET' \
  'http://localhost:8080/api/v1/devices?enterprise_id=acme&site_id=acme-chicago' \
  -H 'accept: application/json' | jq .

curl --fail -X 'GET' \
  'http://localhost:8080/api/v1/sim_cards?enterprise_id=acme&site_id=acme-chicago' \
  -H 'accept: application/json' | jq .

curl --fail -X 'GET' \
  'http://localhost:8080/api/v1/slices?enterprise_id=acme&site_id=acme-chicago' \
  -H 'accept: application/json' | jq .

echo -e "\n--- done testing REST"

echo "Testing GraphQL server"

curl --fail 'http://localhost:8080/graphql' \
    -H 'accept: application/json, multipart/mixed' \
    -H 'content-type: application/json' \
    -d '{ "query":"query{ enterprises {id name description}}"}' \
    | jq .

curl --fail 'http://localhost:8080/graphql' \
    -H 'accept: application/json, multipart/mixed' \
    -H 'content-type: application/json' \
    -d '{ "query":"query{ applications(enterpriseID: \"acme\") {id name}}"}' \
    | jq .

curl --fail 'http://localhost:8080/graphql' \
    -H 'accept: application/json, multipart/mixed' \
    -H 'content-type: application/json' \
    -d '{ "query":"query{ devices(enterpriseID: \"acme\", siteID: \"acme-chicago\") {id name description}}"}' \
    | jq .

curl --fail 'http://localhost:8080/application-query' \
    -H 'accept: application/json, multipart/mixed' \
    -H 'content-type: application/json' \
    --data-raw '{"query":"query {\n  applicationServiceGetApplications(in: {enterpriseId: \"acme\"} ) {\n    applications {\n      iD\n      endpoint {\n        iD\n       displayName\n}\n    }\n  }\n}"}' \
    | jq .

curl --fail 'http://localhost:8080/enterprise-query' \
  -H 'accept: application/json, multipart/mixed' \
  -H 'content-type: application/json' \
  --data-raw '{"query":"query {\n  enterpriseServiceGetEnterprises {\n    enterprises {\n      iD\n    }\n  }\n}","variables":null}' \
  | jq .

echo -e "\n--- done testing GraphQL"