package string

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

func TestSubStr(t *testing.T) {
	str := "hello world"
	Convey("TestSubStr", t, func() {
		Convey("success", func() {
			l, r, err := SubStr(str, 6, 5)
			assert.Equal(t, l, int64(5))
			assert.Equal(t, r, "world")
			assert.Equal(t, err, nil)
		})
	})
}
