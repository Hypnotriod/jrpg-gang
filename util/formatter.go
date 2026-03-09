package util

import (
	"fmt"
	"reflect"
	"strings"
)

func AsCommaSeparatedObjectsSlice(slice any) string {
	values := reflect.ValueOf(slice)
	var sb strings.Builder
	for i := 0; i < values.Len(); i++ {
		sb.WriteString(fmt.Sprintf("{%v}", values.Index(i)))
		if i < values.Len()-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}

func AsCommaSeparatedSlice(slice any) string {
	values := reflect.ValueOf(slice)
	var sb strings.Builder
	for i := 0; i < values.Len(); i++ {
		sb.WriteString(fmt.Sprintf("%v", values.Index(i)))
		if i < values.Len()-1 {
			sb.WriteString(", ")
		}
	}
	return sb.String()
}
