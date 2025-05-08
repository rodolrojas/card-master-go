package config

import (
	errorConfig "card-master/lib/middleware/error"

	"github.com/gofiber/fiber/v2"
)

var App = func () fiber.Config {
	return fiber.Config{
		ErrorHandler: errorConfig.Config,
	}
}