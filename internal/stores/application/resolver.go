/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package application

import v1 "github.com/onosproject/aether-operations-api/gen/go/v1"

func NewApplicationResolver(srv v1.ApplicationServiceServer) *v1.ApplicationServiceResolvers {
	r := v1.ApplicationServiceResolvers{
		Service: srv,
	}
	return &r
}
