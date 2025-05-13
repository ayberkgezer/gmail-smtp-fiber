package handler

import (
	"context"

	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/model/request"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/services"
)

// IEmailHandler e-posta iş mantığını soyutlar
type IEmailHandler interface {
	HandleEmail(ctx context.Context, req request.EmailRequest) error
}

type EmailHandler struct {
	service services.IEmailService
}

func NewEmailHandler(s services.IEmailService) IEmailHandler {
	return &EmailHandler{service: s}
}

// sadece service çağırır
func (h *EmailHandler) HandleEmail(ctx context.Context, req request.EmailRequest) error {
	return h.service.SendMail(ctx, req)
}
