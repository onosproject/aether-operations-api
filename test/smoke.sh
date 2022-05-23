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
  'http://localhost:8080/api/v1/enterprises/acme/applications' \
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
    -d '{ "query": "query { enterprises { enterprises { id name description } } }" }' \
    | jq .

curl --fail 'http://localhost:8080/graphql' \
    -H 'accept: application/json, multipart/mixed' \
    -H 'content-type: application/json' \
    -d '{ "query": "query { applications(enterpriseID: \"acme\") { applications { id name description } } }" }' \
    | jq .

curl --fail 'http://localhost:8080/graphql' \
    -H 'accept: application/json, multipart/mixed' \
    -H 'content-type: application/json' \
    -d '{ "query": "query { sites(enterpriseID: \"acme\") { sites { id name description devices { id name } } } }" }' \
    | jq .

curl --fail 'http://localhost:8080/graphql' \
    -H 'accept: application/json, multipart/mixed' \
    -H 'content-type: application/json' \
    -d '{ "query": "query { slices(enterpriseID: \"acme\", siteID: \"acme-chicago\") { slices { id name description } } }" }' \
    | jq .

echo -e "\n--- done testing GraphQL"
