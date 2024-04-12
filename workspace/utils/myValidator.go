package utils

import "github.com/go-playground/validator/v10"

var Validate *validator.Validate

func InitValidator() {
	Validate = validator.New()
	_ = Validate.RegisterValidation("isInSlice", isInSlice)

}
func isInSlice(fl validator.FieldLevel) bool {
	var slice = []string{"red", "green", "blue"}
	value := fl.Field().String()
	for _, v := range slice {
		if v == value {
			return true
		}
	}
	return false
}
