package assert

import "reflect"

// IsNil check interface is nill
func IsNil(data interface{}) bool {
	if data == nil {
		return true
	}

	kind := reflect.ValueOf(data).Kind()
	if kind == reflect.Slice ||
		kind == reflect.Map ||
		kind == reflect.Chan ||
		kind == reflect.Interface ||
		kind == reflect.Func ||
		kind == reflect.Ptr {
		return reflect.ValueOf(data).IsNil()
	}

	return false
}
