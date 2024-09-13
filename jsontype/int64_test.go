package jsontype

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInt64(t *testing.T) {
	type Data struct {
		Value Int64 `json:"value"`
	}
	d := Data{
		Value: Int64(123),
	}
	b, err := json.Marshal(d)
	assert.Nil(t, err)
	assert.Equal(t, `{"value":123}`, string(b))

	tests := map[string]int64{
		`{"value":123}`:     123,
		`{"value":"123"}`:   123,
		`{"value":123.1}`:   123,
		`{"value":"123.1"}`: 123,
		`{"value":""}`:      0,
		`{"value":null}`:    0,
		`{"value":"null"}`:  0,
		`{"value":false}`:   0,
		`{"value":true}`:    1,
	}
	for str, want := range tests {
		t.Run(str, func(t *testing.T) {
			s := &Data{}
			if err := json.Unmarshal([]byte(str), s); err != nil {
				t.Fatal(err)
			}
			if s.Value.ToInt64() != want {
				t.Fatalf("s.Value=%v, want=%v", s.Value, want)
			}
		})
	}
}
