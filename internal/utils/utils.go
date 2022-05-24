/*
 * SPDX-FileCopyrightText: 2022-present Intel Corporation
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package utils

import (
	"reflect"
)

func PointerToString(p *string) string {
	if reflect.ValueOf(p).IsNil() {
		return ""
	}
	return *p
}

func PointerToInt64(p interface{}) int64 {
	k := reflect.ValueOf(p).Kind()

	if p == nil {
		return 0
	}

	if k == reflect.Ptr {
		if reflect.ValueOf(p).IsNil() {
			return 0
		}
		p = reflect.Indirect(reflect.ValueOf(p)).Interface()
		k = reflect.ValueOf(p).Kind()
	}

	switch k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return reflect.ValueOf(p).Int()
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return int64(reflect.ValueOf(p).Uint())
	default:
		return reflect.ValueOf(p).Int()
	}
}
