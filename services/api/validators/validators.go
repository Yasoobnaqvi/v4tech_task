package validators

import (
	"gopkg.in/go-playground/validator.v9"
)

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       interface{}
}

var validate = validator.New()

func ValidateStruct(postBody interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(postBody)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Value()
			errors = append(errors, &element)
		}
	}
	return errors
}

type AddProductPostBody struct {
	Name  	      string    `json:"name"`
	Description   string    `json:"description"`
	Price      	  float32   `json:"price"`
}
