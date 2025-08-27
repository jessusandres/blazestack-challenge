package helpers

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func GetErrorMsg(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "This field is required"
	case "email":
		return "Should be a valid email"
	case "min":
		return fmt.Sprintf("Should be at least %s characters long", fe.Param())
	case "len":
		return fmt.Sprintf("Should be exactly %s characters long", fe.Param())
	case "max":
		return fmt.Sprintf("Should be at most %s characters long", fe.Param())
	case "gte":
		return fmt.Sprintf("Should be greater than or equal to %s", fe.Param())
	case "lte":
		return fmt.Sprintf("Should be less than or equal to %s", fe.Param())
	}

	return "Unknown error"
}
