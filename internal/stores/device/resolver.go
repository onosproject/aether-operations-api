/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package device

import v1 "github.com/onosproject/aether-operations-api/gen/go/v1"

func NewDeviceResolver(srv v1.DeviceServiceServer) *v1.DeviceServiceResolvers {
	r := v1.DeviceServiceResolvers{
		Service: srv,
	}
	return &r
}
