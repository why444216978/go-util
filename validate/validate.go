package validate

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	util_str "github.com/why444216978/go-util/string"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// Validate example
// type Data struct {
// 	TpPrescriptionCode string   `json:"tp_prescription_code" validate:"required,min=1,max=32"`
// 	PrescriptionType   uint8    `json:"prescription_type" validate:"required,oneof=1 2"`
// 	HospitalCode       string   `json:"hospital_code" validate:"required,min=1,max=32"`
// 	PharmacyCode       string   `json:"pharmacy_code" validate:"required,min=1,max=32"`
// 	OpenedTime         int      `json:"opened_time" validate:"required,min=1000000000,max=1999999999"`
// 	AppCode            string   `json:"app_code" validate:"required,min=1,max=32" `
// 	PatientAge         uint8    `json:"patient_age" validate:"required,min=1,max=200"`
// 	PatientSex         uint8    `json:"patient_sex" validate:"required,oneof=1 2"`
// 	PatientName        string   `json:"patient_name" validate:"required,min=1,max=32"`
// 	PatientPhone       string   `json:"patient_phone" validate:"required,len=11"`
// }
// data := Data{}
// err = Validate(data)
// if err != nil {
// 	fmt.Println(err.Error())
// }
func Validate(val interface{}) error {
	err := validate.Struct(val)
	if err == nil {
		return nil
	}

	for _, err := range err.(validator.ValidationErrors) {
		field := util_str.CamelToSnake(err.Field())
		return errors.New(fmt.Sprintf("param %s must %s %s", field, err.Tag(), err.Param()))

	}
	return nil
}
