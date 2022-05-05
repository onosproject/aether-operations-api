/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package openapiv3

import (
	"embed"
)

//go:embed dist/*
var OpenAPI embed.FS
