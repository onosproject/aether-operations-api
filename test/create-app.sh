#
# SPDX-FileCopyrightText: 2022-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0
#

curl -X 'POST' \
  'http://localhost:8080/api/v1/applications' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "applicationId": "test-app-id",
  "name": "test-app-name",
  "description": "foo-description",
  "address": "bar-address",
  "endpoints": [
    {
      "endpointId": "da",
      "name": "data acquisition endpoint",
      "mbr": {
        "uplink": "2000000",
        "downlink": "1000000"
      },
      "portStart": 7585,
      "portEnd": 7588,
      "protocol": "TCP"
    }
  ],
  "enterpriseId": "acme"
}'