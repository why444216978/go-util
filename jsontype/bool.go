package jsontype

import (
	"bytes"
	"encoding/json"
	"strconv"
)

var (
	_ json.Marshaler   = (*Bool)(nil)
	_ json.Unmarshaler = (*Bool)(nil)
)

type Bool bool

func (v *Bool) UnmarshalJSON(b []byte) error {
	s := string(bytes.Trim(b, `"`))
	switch s {
	case "null":
		*v = false
		return nil
	case "":
		*v = false
		return nil
	}

	r, e := strconv.ParseBool(s)
	if e != nil {
		return e
	}
	*v = Bool(r)
	return nil
}

func (v Bool) MarshalJSON() ([]byte, error) {
	if v.ToBool() {
		return []byte("true"), nil
	}
	return []byte("false"), nil
}

func (v Bool) ToBool() bool {
	return bool(v)
}
