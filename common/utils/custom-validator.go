package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func CustomValidator [T any](structRef *T) fiber.Handler {
	return func (c *fiber.Ctx) error {
		var parsedValue T
		if err := c.BodyParser(&parsedValue); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
		}
		validate := validator.New()
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
									fErr.Message = fmt.Sprintf("%s is required", fe.Field())
							case "oneof":
									allowed := strings.Fields(fe.Param()) // split the space-separated values
									fErr.Allowed = allowed
									fErr.Message = fmt.Sprintf("%s must be one of %v", fe.Field(), allowed)
							default:
									fErr.Message = fmt.Sprintf("%s failed on '%s' validation", fe.Field(), fe.Tag())
							}
							out = append(out, fErr)
					}
					return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
						"error":  "Validation failed",
						"fields": out,
					})
				}
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Validation failed"})
		}
		return c.Next()
	}
}