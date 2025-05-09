package main

import (
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Initialize the server
	server.NewServer(app).StartHttpServer()
}
