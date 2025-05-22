package middleware

import (
	"log"

	"github.com/ayberkgezer/gmail-smtp-fiber/internal/base"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/config"
	"github.com/gofiber/fiber/v2"
)

func APIKeyMiddleware(ctx *fiber.Ctx) error {
	if ctx.Get("X-API-KEY") != config.ApiKey {
		// Log the invalid API key usage
		log.Printf("Invalid API Key:")
		return ctx.Status(fiber.StatusUnauthorized).JSON(base.NewErrorResponse(fiber.StatusUnauthorized, "Invalid API Key"))
	}
	return ctx.Next()
}
