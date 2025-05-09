package main

import (
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/pkg/server"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	server.NewServer(app).StartHttpServer()
}
