/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package devicegroup

import v1 "github.com/onosproject/aether-operations-api/gen/go/v1"

func NewDeviceGroupResolver(srv v1.DeviceGroupServiceServer) *v1.DeviceGroupServiceResolvers {
	r := v1.DeviceGroupServiceResolvers{
		Service: srv,
	}
	return &r
}
