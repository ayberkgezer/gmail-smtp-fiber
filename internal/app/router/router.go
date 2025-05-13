package router

import (
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/controller"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/base"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App, emailController controller.IEmailController) {
	// Initialize the API group
	api := app.Group("/api/v1")

	// GET /api/v1/health
	api.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(base.NewBaseResponse(fiber.StatusOK, "Server Running"))
	})

	// email group
	email := api.Group("/email")
	// GET /api/v1/email/send
	email.Post("/send", emailController.SendEmail)
}
