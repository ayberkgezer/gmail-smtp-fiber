package httpfilter

import (
	"log"

	"github.com/ayberkgezer/gmail-smtp-fiber/internal/base"
	"github.com/gofiber/fiber/v2"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	log.Printf("Error occurred: %v", err)
	return ctx.Status(fiber.StatusInternalServerError).JSON(base.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
}
