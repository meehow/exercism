package flatten

import (
	"reflect"
)

// Flatten returns flattened list with all values except nil/null
func Flatten(input interface{}) []interface{} {
	output := make([]interface{}, 0)
	flattenSlice(input, &output)
	return output
}

func flattenSlice(input interface{}, output *[]interface{}) {
	s := reflect.ValueOf(input)
	for i := 0; i < s.Len(); i++ {
		el := s.Index(i)
		if el.IsNil() {
			continue
		}
		val := el.Interface()
		if reflect.TypeOf(val).Kind() == reflect.Slice {
			flattenSlice(val, output)
		} else {
			*output = append(*output, val)
		}
	}
}
