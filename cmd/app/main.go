package main

import (
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/router"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.InitializeRoutes(app)
	// Initialize the server
	server.NewServer(app).StartHttpServer()
}
