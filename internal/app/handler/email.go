package handler

import (
	"context"

	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/services"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/domain"
	"github.com/ayberkgezer/gocolorlog"
)

type IEmailHandler interface {
	HandleEmail(ctx context.Context, email *domain.Email, reqID string) error
}

type EmailHandler struct {
	service services.IEmailService
}

func NewEmailHandler(s services.IEmailService) IEmailHandler {
	return &EmailHandler{service: s}
}

func (h *EmailHandler) HandleEmail(ctx context.Context, email *domain.Email, reqID string) error {
	defer gocolorlog.Infof("EmailHandler.HandleEmail RequestID:%s", reqID)
	return h.service.SendMail(ctx, email, reqID)
}
