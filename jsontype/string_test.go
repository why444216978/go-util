package jsontype

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestString(t *testing.T) {
	type Data struct {
		Value String `json:"value"`
	}

	d := Data{
		Value: "123",
	}
	b, err := json.Marshal(d)
	assert.Nil(t, err)
	assert.Equal(t, `{"value":"123"}`, string(b))

	tests := map[string]string{
		`{"value":"\"s\"\n"}`: "\"s\"\n",
		`{"value":123}`:       "123",
		`{"value":123.4}`:     "123.4",
		`{"value":-1}`:        "-1",
		`{"value":0}`:         "0",
		`{"value":"1"}`:       "1",
		`{"value":false}`:     "false",
		`{"value":true}`:      "true",
		`{"value":null}`:      "",
		`{"value":""}`:        "",
	}
	for str, want := range tests {
		t.Run(str, func(t *testing.T) {
			s := &Data{}
			if err := json.Unmarshal([]byte(str), s); err != nil {
				t.Fatal(err)
			}
			if s.Value.ToString() != want {
				t.Fatalf("s.Value=%v, want=%v", s.Value, want)
			}
		})
	}
}
