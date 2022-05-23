/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package enterprise

import (
	v1 "github.com/onosproject/scaling-umbrella/gen/go/v1"
)

func NewEnterpriseResolver(srv v1.EnterpriseServiceServer) *v1.EnterpriseServiceResolvers {
	r := v1.EnterpriseServiceResolvers{
		Service: srv,
	}
	return &r
}
