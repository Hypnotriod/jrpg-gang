package util

import (
	"fmt"
	"reflect"
	"strings"
)

func AsCommaSeparatedSlice(slice interface{}) string {
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
