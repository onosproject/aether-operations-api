/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

import (
	"flag"
	"strings"
	"time"
)

type arrayFlags []string

func (i *arrayFlags) String() string {
	s := []string(*i)
	return strings.Join(s, ", ")
}

func (i *arrayFlags) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type DataSourcesConfig struct {
	OnosConfigAddress string
	OnosConfigTimeout time.Duration
	OnosTopoAddress   string
}
type ServersConfig struct {
	GrpcAddress string
	HttpAddress string
	Cors        arrayFlags
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
	LogLevel      string
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

	flag.StringVar(&config.LogLevel, logLevelParam, logLevel, "Log Level")
	flag.StringVar(&config.DataSources.OnosConfigAddress, onosConfigAddressParam, defaultOnosConfigAddress, "The ONOS Config address")
	flag.DurationVar(&config.DataSources.OnosConfigTimeout, onosConfigTimeoutParam, defaultOnosConfigTimeout, "The ONOS Config timeout")
	flag.StringVar(&config.DataSources.OnosTopoAddress, onosTopoAddressParam, defaultOnosTopoAddress, "The ONOS Topo address")
	flag.StringVar(&config.ServersConfig.GrpcAddress, grpcServerAddressParam, defaultGrpcAddress, "The gRPC Server address")
	flag.StringVar(&config.ServersConfig.HttpAddress, httpServerAddressParam, defaultHttpAddress, "The HTTP Server address")
	flag.Var(&config.ServersConfig.Cors, corsOriginParam, "CORS origin (repeated)")

	flag.Parse()

	return config
}
