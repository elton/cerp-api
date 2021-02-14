package controllers

import (
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/elton/cerp-api/utils/logger"
	"github.com/gofiber/fiber/v2"
)

// Server represents the struct of our server.
type Server struct {
	Router *fiber.App
}

// Run API Server endpoint.
func (s *Server) Run(port string) {
	// Logging
	date := time.Now().Format("20060102")
	f, err := os.Create("access-cerp-" + date + ".log")
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}()

	// Web server
	// Custom config
	app := fiber.New(fiber.Config{
		Prefork:       true,
		StrictRouting: true,
		ServerHeader:  "CERP-API-Server",
		ReadTimeout:   10 * time.Second,
		WriteTimeout:  10 * time.Second,
	})

	s.initializeRouters(app)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		_ = <-c
		logger.Info.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(port); err != nil {
		log.Panic(err)
	}

	logger.Info.Printf("Listening to port %s\n", port)
}
