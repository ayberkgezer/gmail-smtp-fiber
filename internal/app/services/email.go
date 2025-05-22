package services

import (
	"context"
	"fmt"

	"github.com/ayberkgezer/gmail-smtp-fiber/internal/domain"
	"github.com/ayberkgezer/gocolorlog"
)

// IEmailService e-posta gönderme işini soyutlar
type IEmailService interface {
	SendMail(ctx context.Context, email *domain.Email, reqID string) error
}

// emailService, IEmailSender'ı kullanarak mail yollar
type emailService struct {
	sender IEmailSender
}

// NewEmailService, SMTP sender’ı enjekte ederek bir IEmailService döner
func NewEmailService(sender IEmailSender) IEmailService {
	return &emailService{sender: sender}
}

// SendMail, request içindeki alanları alıp SMTP sender’a iletir
func (s *emailService) SendMail(ctx context.Context, email *domain.Email, reqID string) error {
	defer gocolorlog.Infof("emailService.SendMail RequestId:%s", reqID)
	subject := fmt.Sprintf("From:%s", email.SenderName)
	emailBody := fmt.Sprintf("From:%s \nMessage:%s", email.SenderEmail, email.Body)
	return s.sender.Send(ctx, subject, emailBody, reqID)
}
