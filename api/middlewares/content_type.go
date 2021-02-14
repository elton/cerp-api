package middlewares

import (
	"github.com/gofiber/fiber/v2"
)

// SetMiddlewareJSON a middleware for JSON responses
func SetMiddlewareJSON() fiber.Handler {
	return func(c *fiber.Ctx) error {
		c.Type("json", "utf-8") // => "application/json; charset=utf-8"
		c.Next()
		return nil
	}
}
