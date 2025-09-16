package utils

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Global validator instance with custom configurations
var validate *validator.Validate

func init() {
	validate = validator.New()
	
	// Register function to get tag name from json tag for better field names
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func CustomValidator[T any](structRef *T) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var parsedValue T
		if err := c.BodyParser(&parsedValue); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid JSON format",
				"details": err.Error(),
			})
		}
		
		err := validate.Struct(&parsedValue)
		if err != nil {
			if verrs, ok := err.(validator.ValidationErrors); ok {
				type FieldError struct {
					Field    string   `json:"field"`
					Tag      string   `json:"tag"`
					Value    any      `json:"value"`
					Allowed  []string `json:"allowed,omitempty"`
					Message  string   `json:"message"`
				}
				out := make([]FieldError, 0, len(verrs))
				for _, fe := range verrs {
					fErr := FieldError{
						Field: fe.Field(),
						Tag:   fe.Tag(),
						Value: fe.Value(),
					}
					
					switch fe.Tag() {
					case "required":
						fErr.Message = fmt.Sprintf("The field '%s' is required", fe.Field())
					case "oneof":
						allowed := strings.Fields(fe.Param())
						fErr.Allowed = allowed
						fErr.Message = fmt.Sprintf("The field '%s' must be one of: %s", fe.Field(), strings.Join(allowed, ", "))
					case "min":
						fErr.Message = fmt.Sprintf("The field '%s' must be at least %s characters long", fe.Field(), fe.Param())
					case "max":
						fErr.Message = fmt.Sprintf("The field '%s' must be at most %s characters long", fe.Field(), fe.Param())
					case "email":
						fErr.Message = fmt.Sprintf("The field '%s' must be a valid email address", fe.Field())
					case "url":
						fErr.Message = fmt.Sprintf("The field '%s' must be a valid URL", fe.Field())
					case "numeric":
						fErr.Message = fmt.Sprintf("The field '%s' must be a number", fe.Field())
					case "gt":
						fErr.Message = fmt.Sprintf("The field '%s' must be greater than %s", fe.Field(), fe.Param())
					case "gte":
						fErr.Message = fmt.Sprintf("The field '%s' must be greater than or equal to %s", fe.Field(), fe.Param())
					case "lt":
						fErr.Message = fmt.Sprintf("The field '%s' must be less than %s", fe.Field(), fe.Param())
					case "lte":
						fErr.Message = fmt.Sprintf("The field '%s' must be less than or equal to %s", fe.Field(), fe.Param())
					default:
						fErr.Message = fmt.Sprintf("The field '%s' failed validation for '%s'", fe.Field(), fe.Tag())
					}
					out = append(out, fErr)
				}
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"error":   "Validation failed",
					"message": "Please check the following fields and try again",
					"fields":  out,
				})
			}
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Validation failed",
				"details": err.Error(),
			})
		}
		
		return c.Next()
	}
}