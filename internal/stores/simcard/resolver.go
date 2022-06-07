/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package simcard

import v1 "github.com/onosproject/aether-operations-api/gen/go/v1"

func NewSimCardResolver(srv v1.SimCardServiceServer) *v1.SimCardServiceResolvers {
	r := v1.SimCardServiceResolvers{
		Service: srv,
	}
	return &r
}
