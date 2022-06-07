/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package smallcell

import v1 "github.com/onosproject/aether-operations-api/gen/go/v1"

func NewSmallCellResolver(srv v1.SmallCellServiceServer) *v1.SmallCellServiceResolvers {
	r := v1.SmallCellServiceResolvers{
		Service: srv,
	}
	return &r
}
