package validate

import (
	"reflect"

	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
)

var validate *validator.Validate

var (
	// customDataTag is default data tag name
	customDataTag = "json"
	// customErrTag is default custom tag name
	customErrTag = "custom_err"
)

func init() {
	validate = validator.New()
}

// SetCustomDataTag set custom data tag name
func SetCustomDataTag(tag string) {
	customDataTag = tag
}

// SetCustomErrTag set custom err tag name
func SetCustomErrTag(tag string) {
	customErrTag = tag
}

// Validate is validate a struct exposed fields
func Validate(val interface{}) error {
	err := validate.Struct(val)
	if err == nil {
		return nil
	}

	for _, err := range err.(validator.ValidationErrors) {
		return wrapErr(val, err)
	}

	return nil
}

// wrapErr is wrap err
func wrapErr(val interface{}, err validator.FieldError) error {
	t := reflect.TypeOf(val)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	f, ok := t.FieldByName(err.Field())
	if !ok {
		return errors.Errorf("param %s must %s %s", err.Field(), err.Tag(), err.Param())
	}

	errTag := f.Tag.Get(customErrTag)
	if errTag == "" {
		return errors.Errorf("param %s must %s %s", f.Tag.Get(customDataTag), err.Tag(), err.Param())
	}

	return errors.Errorf("%s:%s", f.Tag.Get(customDataTag), errTag)
}
