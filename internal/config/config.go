/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import "flag"

type DataSourcesConfig struct {
	OnosConfigAddress string
	OnosTopoAddress   string
}
type ServersConfig struct {
	GrpcAddress string
	HttpAddress string
}

type BuildConfig struct {
	BuildTime  string
	CommitHash string
	VcsDirty   string
	Version    string
}

type Config struct {
	DataSources   *DataSourcesConfig
	ServersConfig *ServersConfig
	BuildConfig   *BuildConfig
}

func GetConfig() *Config {

	config := &Config{
		DataSources:   &DataSourcesConfig{},
		ServersConfig: &ServersConfig{},
		BuildConfig: &BuildConfig{
			BuildTime:  buildTime,
			CommitHash: commitHash,
			VcsDirty:   vcsDirty,
			Version:    version,
		},
	}

	flag.StringVar(&config.DataSources.OnosConfigAddress, onosConfigAddressParam, defaultOnosConfigAddress, "The ONOS Config address")
	flag.StringVar(&config.DataSources.OnosTopoAddress, onosTopoAddressParam, defaultOnosTopoAddress, "The ONOS Topo address")
	flag.StringVar(&config.ServersConfig.GrpcAddress, grpcServerAddressParam, defaultGrpcAddress, "The gRPC Server address")
	flag.StringVar(&config.ServersConfig.HttpAddress, httpServerAddressParam, defaultHttpAddress, "The HTTP Server address")

	flag.Parse()

	return config
}
