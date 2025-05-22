package request

import "github.com/ayberkgezer/gmail-smtp-fiber/internal/domain"

type EmailRequest struct {
	Name         string `json:"name" validate:"required,min=3,max=100"`
	Email        string `json:"email" validate:"required,email"`
	EmailMessage string `json:"emailMessage" validate:"required,min=10,max=1000"`
}

func (r *EmailRequest) ToDomain() *domain.Email {
	return &domain.Email{
		SenderName:  r.Name,
		SenderEmail: r.Email,
		Body:        r.EmailMessage,
	}
}
