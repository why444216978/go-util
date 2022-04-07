package validate

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

func TestValidate(t *testing.T) {
	Convey("TestValidateCamel", t, func() {
		Convey("snake", func() {
			type Data struct {
				CodeA uint8  `json:"code_a" validate:"required,oneof=1 2" custom_err:"必须为1或2"`
				NameA string `json:"name_a" validate:"required" custom_err:"不能为空"`
			}
			data := &Data{
				CodeA: 0,
				NameA: "name",
			}
			err := Validate(data)
			assert.Equal(t, err.Error(), "code_a:必须为1或2")
		})
		Convey("camel", func() {
			type Data struct {
				CodeA uint8  `json:"codeA" validate:"required,oneof=1 2" custom_err:"必须为1或2"`
				NameA string `json:"nameA" validate:"required" custom_err:"不能为空"`
			}
			data := &Data{
				CodeA: 1,
				NameA: "",
			}
			err := Validate(data)
			assert.Equal(t, err.Error(), "nameA:不能为空")
		})
		Convey("upper", func() {
			type Data struct {
				CodeA uint8  `json:"CODEA" validate:"required,oneof=1 2" custom_err:"必须为1或2"`
				NameA string `json:"NAMEA" validate:"required" custom_err:"不能为空"`
			}
			data := &Data{
				CodeA: 1,
				NameA: "",
			}
			err := Validate(data)
			assert.Equal(t, err.Error(), "NAMEA:不能为空")
		})
		Convey("lower", func() {
			type Data struct {
				CodeA uint8  `json:"codea" validate:"required,oneof=1 2" custom_err:"必须为1或2"`
				NameA string `json:"namea" validate:"required" custom_err:"不能为空"`
			}
			data := &Data{
				CodeA: 1,
				NameA: "",
			}
			err := Validate(data)
			assert.Equal(t, err.Error(), "namea:不能为空")
		})
	})
}

func TestSetCustomDataTag(t *testing.T) {
	Convey("TestSetCustomDataTag", t, func() {
		target := "tag"
		SetCustomDataTag(target)
		assert.Equal(t, customDataTag, target)
	})
}

func TestSetCustomErrTag(t *testing.T) {
	Convey("TestSetCustomErrTag", t, func() {
		target := "tag"
		SetCustomErrTag(target)
		assert.Equal(t, customErrTag, target)
	})
}
