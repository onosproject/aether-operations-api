#
# SPDX-FileCopyrightText: $today.year-present Intel Corporation
#
# SPDX-License-Identifier: Apache-2.0
#

FROM golang:1.18.1-stretch as builder

ENV CGO_ENABLED=1
ARG org_label_schema_version=unknown
ARG org_label_schema_vcs_url=unknown
ARG org_label_schema_vcs_ref=unknown
ARG org_label_schema_build_date=unknown
ARG org_opencord_vcs_commit_date=unknown
ARG org_opencord_vcs_dirty=unknown

WORKDIR /app

COPY . ./

RUN go build -mod vendor \
      -ldflags "-w -X github.com/onosproject/aether-operations-api/internal/config.buildTime=${org_label_schema_build_date} \
        -X github.com/onosproject/aether-operations-api/internal/config.commitHash=${org_label_schema_vcs_ref} \
        -X github.com/onosproject/aether-operations-api/internal/config.vcsDirty=${org_opencord_vcs_dirty} \
        -X github.com/onosproject/aether-operations-api/internal/config.version=${org_label_schema_version}" \
      ./cmd/roc-api
RUN ls /app

FROM golang:1.18.1-stretch

WORKDIR /app
COPY --from=builder /app/roc-api /app/roc-api
RUN chmod a+x /app/roc-api
RUN ls /app
CMD [ "/app/roc-api" ]

# Label image
LABEL org.label-schema.schema-version=1.0 \
      org.label-schema.name=scaling-umbrella \
      org.label-schema.version=$org_label_schema_version \
      org.label-schema.vcs-url=$org_label_schema_vcs_url \
      org.label-schema.vcs-ref=$org_label_schema_vcs_ref \
      org.label-schema.build-date=$org_label_schema_build_date \
      org.opencord.vcs-commit-date=$org_opencord_vcs_commit_date \
      org.opencord.vcs-dirty=$org_opencord_vcs_dirty