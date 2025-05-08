package request

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type FieldError struct {
	Field 	string `json:"field"`
	Tag   	string `json:"tag"`
	Message string `json:"message"`
}

type ValidationError struct {
	Fields []FieldError `json:"errors"`
}

func (ve ValidationError) Error() string {
	return "Validation Failed"
}

var validate = validator.New()

func Validate(c *fiber.Ctx, input interface{}) error {
	if err := c.BodyParser(input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON format")
	}
	if err := ValidateInput(input); err != nil {
		jsonErrors, err := json.Marshal(err)
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "Failed to marshal error response")
		}
		return fiber.NewError(fiber.StatusUnprocessableEntity, string(jsonErrors))
	}

	return nil
}

func ValidateInput(input interface{}) error {
	if err := validate.Struct(input); err != nil {
		var validationErrors []FieldError
		for _, e := range err.(validator.ValidationErrors) {
			fmt.Println(e.Field(), e.Tag())
			validationErrors = append(validationErrors, FieldError{
				Field:  e.Field(),
				Tag:    e.Tag(),
				Message: translateError(e.Field(), e.Tag()),
			})
		}		
		return ValidationError{Fields: validationErrors}
	}
	return nil
}

func translateError(field, tag string) string {
	switch tag {
	case "required":
		return fmt.Sprintf("%s is required", field)
	case "min":
		return fmt.Sprintf("%s is too short", field)
	case "max":
		return fmt.Sprintf("%s is too long", field)
	case "email":
		return fmt.Sprintf("%s must be a valid email", field)
	case "gte":
		return fmt.Sprintf("%s must be greater than or equal to minimum", field)
	case "lte":
		return fmt.Sprintf("%s must be less than or equal to maximum", field)
	default:
		return fmt.Sprintf("%s is not valid (%s)", field, tag)
	}
}