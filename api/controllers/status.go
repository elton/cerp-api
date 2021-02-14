package controllers

import (
	"github.com/gofiber/fiber/v2"
)

// HealthCheck health checks
func HealthCheck(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "UP"})
}
