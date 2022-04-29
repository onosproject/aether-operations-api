/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

type DataSourcesConfig struct {
	OnosConfigAddress string
	OnosTopoAddress   string
}
type ServersConfig struct {
	GrpcAddress string
	RestAddress string
	// deprecated
	GraphQlAddress string
}

type Config struct {
	DataSources   *DataSourcesConfig
	ServersConfig *ServersConfig
}

func GetConfig() *Config {
	// TODO add CLI Params

	return &Config{
		DataSources: &DataSourcesConfig{
			OnosConfigAddress: "0.0.0.0:5150",
			OnosTopoAddress:   "0.0.0.0:5151",
		},
		ServersConfig: &ServersConfig{
			GrpcAddress:    "0.0.0.0:50060",
			RestAddress:    "0.0.0.0:8080",
			GraphQlAddress: "0.0.0.0:8081", // FIXME use the same mux server for REST and GraphQL
		},
	}
}
