package error

import "github.com/gofiber/fiber/v2"

type ErrorMessage interface{}

var Config = func (c *fiber.Ctx, err error) error {
	var message ErrorMessage

	// Default status code is 500
	code := fiber.StatusInternalServerError
	message = "Internal Server Error"

	// Check if the error is a fiber.Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
		message = e.Message
		// Try to parse the message as JSON
		var jsonMessage map[string]interface{}
		if msgStr, ok := message.(string); ok {
			if err := c.App().Config().JSONDecoder([]byte(msgStr), &jsonMessage); err == nil {
				message = jsonMessage
			}
		} else{
			message = e.Message
		}
	}

	return c.Status(code).JSON(fiber.Map{
		"error": true,
		"message": message,
		"status": code,
	})
}