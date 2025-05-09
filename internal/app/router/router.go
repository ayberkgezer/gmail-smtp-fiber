package router

import (
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/base"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	api := app.Group("/api/v1")

	// GET /api/v1/health
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(base.NewBaseResponse(fiber.StatusOK, "Server Running"))
	})
	// GET /api/v1/email/send
	//api.Get("/email/send")
}
