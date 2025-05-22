package services

import (
	"context"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/ayberkgezer/gocolorlog"
)

type IEmailSender interface {
	Send(ctx context.Context, subject, body, reqID string) error
}

type smtpSender struct {
	host     string
	port     string
	username string
	password string
	to       string
}

func NewSMTPSender(host, port, user, pass, to string) IEmailSender {
	return &smtpSender{
		host:     host,
		port:     port,
		username: user,
		password: pass,
		to:       to,
	}
}

func (s *smtpSender) Send(ctx context.Context, subject, body, reqID string) error {
	defer gocolorlog.Infof("smtpSender.Send | RequestId:%s", reqID)
	addr := fmt.Sprintf("%s:%s", s.host, s.port)
	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	from := s.username
	to := []string{s.to}
	msg := []byte(
		"From: " + from + "\r\n" +
			"To: " + strings.Join(to, ",") + "\r\n" +
			"Subject: " + subject + "\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/plain; charset=\"utf-8\"\r\n" +
			"\r\n" +
			body + "\r\n")

	return smtp.SendMail(addr, auth, from, to, msg)
}
