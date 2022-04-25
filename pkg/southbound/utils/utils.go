package utils

import "reflect"

func PointerToString(p *string) string {
	if reflect.ValueOf(p).IsNil() {
		return ""
	}
	return *p
}

func PointerToInt64(p *int64) int64 {
	if reflect.ValueOf(p).IsNil() {
		return 0
	}
	return *p
}