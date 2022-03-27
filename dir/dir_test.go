package dir

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

func TestGetFileInfo(t *testing.T) {
	Convey("TestGetFileInfo", t, func() {
		Convey("success", func() {
			f := "/user/local/data.toml"
			info, err := GetPathInfo(f)
			assert.Equal(t, info.Path, f)
			assert.Equal(t, info.Base, "data.toml")
			assert.Equal(t, info.BaseNoExt, "data")
			assert.Equal(t, info.Ext, ".toml")
			assert.Equal(t, info.ExtNoSpot, "toml")
			assert.Equal(t, err, nil)
		})
	})
}
