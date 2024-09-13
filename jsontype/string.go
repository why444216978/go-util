package jsontype

import (
	"bytes"
	"encoding/json"
)

var (
	_ json.Marshaler   = (*String)(nil)
	_ json.Unmarshaler = (*String)(nil)
)

type String string

func (v *String) UnmarshalJSON(b []byte) (e error) {
	// 0x22=="
	if len(b) > 1 && b[0] == 0x22 && b[len(b)-1] == 0x22 {
		var s string
		if err := json.Unmarshal(b, &s); err != nil {
			return err
		}
		*v = String(s)
		return nil
	}
	if bytes.Equal(b, []byte("null")) {
		*v = ""
		return nil
	}
	*v = String(b)
	return nil
}

func (v String) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(v))
}

func (a String) ToString() string {
	return string(a)
}
