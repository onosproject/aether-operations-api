/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

// Default parameters name
const logLevelParam = "log-level"
const onosConfigAddressParam = "onos-config-address"
const onosTopoAddressParam = "onos-topo-address"
const grpcServerAddressParam = "grpc-server"
const httpServerAddressParam = "http-server"
const corsOriginParam = "cors-origin"

// Default values
const logLevel = "info"

// Default values for southbound resources
const defaultOnosConfigAddress = "0.0.0.0:5150"
const defaultOnosTopoAddress = "0.0.0.0:5151"

// Defaults for exposed (northbound) resources
const defaultGrpcAddress = "0.0.0.0:50060"
const defaultHttpAddress = "0.0.0.0:8080"

// Variables that are overridden at build time
var buildTime string
var commitHash string
var vcsDirty string
var version string
