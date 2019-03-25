package help

import (
	"reflect"
)

func InterfaceToSlice(arr interface{}) []interface{} {
	v := reflect.ValueOf(arr)

	if v.Kind() != reflect.Slice {
		return make([]interface{}, 0)
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret
}
