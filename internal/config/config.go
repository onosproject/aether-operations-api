/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package config

type Config struct {
	OnosConfigAddress string
	OnosTopoAddress   string
}

func GetConfig() *Config {
	// TODO add CLI Params

	return &Config{
		OnosConfigAddress: "localhost:5150",
		OnosTopoAddress:   "localhost:5151",
	}
}
