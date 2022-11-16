package utils

import (
	"github.com/go-playground/validator/v10"
)

func IsRequestValid(request interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return false, err
	}
	return true, nil

}
