package controllers

import (
	"github.com/elton/cerp-api/api/middlewares"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (s *Server) initializeRouters(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New(), middlewares.SetMiddlewareJSON())
	api.Get("/status", HealthCheck)

	v1 := api.Group("/v1")
	{
		v1.Get("/shops", s.GetAllShops)
		v1.Get("/amount", s.GetAmountByShop)
	}
}
