package main

import (
	"fmt"

	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/controller"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/handler"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/middleware"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/router"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/services"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/common/httpfilter"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/config"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/server"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover" // Fiber'ın recover middleware'i import edildi
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: httpfilter.ErrorHandler,
	})

	app.Use(recover.New())

	app.Use(middleware.Logger)
	app.Use(middleware.APIKeyMiddleware)

	smtpSender := services.NewSMTPSender(
		config.SMTPHost,
		config.SMTPPort,
		config.SMTPUsername,
		config.SMTPPassword,
		config.EMailFrom,
	)

	emailService := services.NewEmailService(smtpSender)
	emailHandler := handler.NewEmailHandler(emailService)
	emailController := controller.NewEmailController(emailHandler)
	// Initialize the router
	router.InitializeRoutes(app, emailController)

	fmt.Printf("Starting server on port %s in %s mode\n", config.Port, config.ENV)
	// Initialize the server
	server.NewServer(app).StartHttpServer()
}
