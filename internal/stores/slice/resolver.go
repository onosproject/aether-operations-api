/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package slice

import v1 "github.com/onosproject/aether-operations-api/gen/go/v1"

func NewSliceResolver(srv v1.SliceServiceServer) *v1.SliceServiceResolvers {
	r := v1.SliceServiceResolvers{
		Service: srv,
	}
	return &r
}
