package util

import (
	"reflect"
	"runtime"
	"strings"
)

func GetFunctionFullName(i any) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

func GetFunctionPathName(i any) (path, name string) {
	full := GetFunctionFullName(i)
	lastSlash := strings.LastIndexByte(full, '/')
	if lastSlash < 0 {
		lastSlash = 0
	}
	lastDot := strings.LastIndexByte(full[lastSlash:], '.') + lastSlash

	return full[:lastDot], full[lastDot+1:]
}

func GetTypePathName(i any) (path, name string) {
	return reflect.TypeOf(i).Elem().PkgPath(), reflect.TypeOf(i).Elem().Name()
}
