package jsontype

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBool(t *testing.T) {
	type Data struct {
		Value Bool `json:"value"`
	}
	d := Data{
		Value: Bool(true),
	}
	b, err := json.Marshal(d)
	assert.Nil(t, err)
	assert.Equal(t, `{"value":true}`, string(b))

	tests := map[string]bool{
		`{"value":false}`:   false,
		`{"value":true}`:    true,
		`{"value":"false"}`: false,
		`{"value":"true"}`:  true,
		`{"value":"null"}`:  false,
		`{"value":""}`:      false,
		`{"value":"1"}`:     true,
		`{"value":"0"}`:     false,
		`{"value":1}`:       true,
		`{"value":0}`:       false,
	}
	for str, want := range tests {
		t.Run(str, func(t *testing.T) {
			s := &Data{}
			if err := json.Unmarshal([]byte(str), s); err != nil {
				t.Fatal(err)
			}
			if s.Value.ToBool() != want {
				t.Fatalf("s.Value=%v, want=%v", s.Value, want)
			}
		})
	}
}
