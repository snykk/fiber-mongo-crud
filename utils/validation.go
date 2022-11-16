package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"

	"github.com/snykk/fiber-mongo-crud/constant"
)

func IsRequestValid(request interface{}) (bool, error) {
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		return false, err
	}
	return true, nil

}

func IsPriorityValid(priority string) error {
	if !isArrayContains(constant.ListOfPriority, priority) {
		var option string
		for index, priority := range constant.ListOfPriority {
			option += priority
			if index != len(constant.ListOfPriority)-1 {
				option += ", "
			}
		}

		return fmt.Errorf("priority must be one of [%s]", option)
	}

	return nil
}

func isArrayContains(arr []string, str string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}
