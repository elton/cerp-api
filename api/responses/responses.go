package responses

import (
	"github.com/gofiber/fiber/v2"
)

// ResultJSON show the result of responses using JSON.
func ResultJSON(c *fiber.Ctx, statusCode int, data interface{}, err error) {
	if err != nil {
		c.JSON(fiber.Map{"status": statusCode, "data": nil, "error": err.Error()})
		return
	}
	c.JSON(fiber.Map{"status": statusCode, "data": data, "error": nil})
}
