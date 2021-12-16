package md5

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

func TestMD5(t *testing.T) {
	Convey("TestMD5", t, func() {
		Convey("success", func() {
			res := MD5("a")
			assert.Equal(t, len(res), 32)
		})
	})
}
