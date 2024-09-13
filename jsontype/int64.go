package jsontype

import (
	"bytes"
	"encoding/json"
	"reflect"
	"strconv"
	"strings"
)

// 123->123
// "123"->123
// true->1 false -> 0
// null->0
// "123.1"->123
// 123.1->123
// ""->0
type Int64 int64

var (
	_ json.Marshaler   = (*Int64)(nil)
	_ json.Unmarshaler = (*Int64)(nil)
)

func (v *Int64) UnmarshalJSON(b []byte) error {
	s := string(bytes.Trim(b, `"`))
	arr := strings.Split(s, ".")
	switch len(arr) {
	case 2:
		fallthrough
	case 1:
		// int 类型
		n, err := strconv.ParseInt(arr[0], 10, 64)
		if err == nil {
			*v = Int64(n)
			return nil
		}
	}
	switch s {
	case "false", "null", "":
		*v = 0
		return nil
	case "true":
		*v = 1
		return nil
	}
	return &json.InvalidUnmarshalError{Type: reflect.TypeOf(v)}
}

func (v Int64) MarshalJSON() ([]byte, error) {
	s := strconv.FormatInt(int64(v), 10)
	return []byte(s), nil
}

func (v Int64) ToInt64() int64 {
	return int64(v)
}
