/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package site

import v1 "github.com/onosproject/scaling-umbrella/gen/go/v1"

func NewSiteResolver(srv v1.SiteServiceServer) *v1.SiteServiceResolvers {
	r := v1.SiteServiceResolvers{
		Service: srv,
	}
	return &r
}
