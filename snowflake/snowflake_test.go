package snowflake

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

func TestGenerate(t *testing.T) {
	Convey("TestGenerate", t, func() {
		Convey("success", func() {
			assert.Equal(t, Generate() > 0, true)
		})
	})
}
