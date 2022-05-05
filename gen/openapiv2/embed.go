/*
 * SPDX-FileCopyrightText: $today.year-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package openapiv2

import (
	"embed"
)

//go:embed dist/*
var OpenAPI embed.FS
