package controller

import (
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/handler"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/model/request"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/base"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/common/middleware"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/common/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

// IEmailController HTTP katmanını soyutlar
type IEmailController interface {
	SendEmail(c *fiber.Ctx) error
}

type EmailController struct {
	handler handler.IEmailHandler
}

func NewEmailController(h handler.IEmailHandler) IEmailController {
	return &EmailController{handler: h}
}

func (ctr *EmailController) SendEmail(c *fiber.Ctx) error {
	var req request.EmailRequest

	// JSON parse
	if err := c.BodyParser(&req); err != nil {
		errMsg := "Invalid request body"
		return c.Status(fiber.StatusBadRequest).
			JSON(base.NewErrorResponse(fiber.StatusBadRequest, errMsg))
	}

	//validation
	if err := validation.ValidateStruct(&req); err != nil {
		return err
	}

	//reqid
	reqID := middleware.GetRequestID(c)
	//
	email := req.ToDomain()

	if err := ctr.handler.HandleEmail(c.Context(), email, reqID); err != nil {
		return err
	}
	return c.Status(fiber.StatusOK).
		JSON(base.NewBaseResponse(fiber.StatusOK, "Email successfully sent"))
}
