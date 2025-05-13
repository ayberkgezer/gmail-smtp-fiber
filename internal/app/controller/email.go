package controller

import (
	"fmt"
	"strings"

	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/handler"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/model/request"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/base"
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
		errMsg := "Geçersiz JSON"
		return c.Status(fiber.StatusBadRequest).
			JSON(base.NewErrorResponse(fiber.StatusBadRequest, errMsg))
	}
	// validator/v10 ile struct validasyonu
	if err := validate.Struct(&req); err != nil {
		verrs := err.(validator.ValidationErrors)
		msgs := make([]string, len(verrs))
		for i, e := range verrs {
			msgs[i] = fmt.Sprintf("%s geçersiz (tag=%s, param=%s)", e.Field(), e.ActualTag(), e.Param())
		}
		return c.Status(fiber.StatusBadRequest).
			JSON(base.NewErrorResponse(fiber.StatusBadRequest, strings.Join(msgs, ", ")))
	}

	if err := ctr.handler.HandleEmail(c.Context(), req); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(base.NewErrorResponse(fiber.StatusInternalServerError, err.Error()))
	}
	return c.Status(fiber.StatusOK).
		JSON(base.NewBaseResponse(fiber.StatusOK, "E-posta başarıyla gönderildi"))
}
