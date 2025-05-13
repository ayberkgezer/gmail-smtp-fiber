package services

import (
	"context"

	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/model/request"
)

// IEmailService e-posta gönderme işini soyutlar
type IEmailService interface {
	SendMail(ctx context.Context, req request.EmailRequest) error
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
func (s *emailService) SendMail(ctx context.Context, req request.EmailRequest) error {
	subject := req.Name
	body := req.EmailMessage
	return s.sender.Send(ctx, req.Email, subject, body)
}
