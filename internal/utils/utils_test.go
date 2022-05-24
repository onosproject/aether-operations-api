/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointerToInt64(t *testing.T) {

	v1 := int64(12)
	v2 := uint64(12)
	var v3 *uint64 = nil

	tests := []struct {
		name string
		args interface{}
		want int64
	}{
		{"int64", v1, v1},
		{"int64-ptr", &v1, v1},
		{"uint64", v2, v1},
		{"uint64-ptr", &v2, v1},
		{"nil", nil, 0},
		{"nil-ptr", v3, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := PointerToInt64(tt.args)
			assert.Equal(t, tt.want, res)
		})
	}
}
