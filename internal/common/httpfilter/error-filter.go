package httpfilter

import (
	"errors"
	"os"

	"github.com/ayberkgezer/gmail-smtp-fiber/internal/base"
	"github.com/ayberkgezer/gmail-smtp-fiber/internal/common/middleware"
	"github.com/ayberkgezer/gocolorlog"
	"github.com/gofiber/fiber/v2"
)

type ValidationError struct {
	Msg string
}

func (v *ValidationError) Error() string {
	return v.Msg
}

func NewValidationError(msg string) error {
	return &ValidationError{Msg: msg}
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {

	env := os.Getenv("ENV")
	showDetails := env == "development"

	status := fiber.StatusInternalServerError
	method := ctx.Method()
	path := ctx.OriginalURL()
	ip := ctx.IP()
	reqID := middleware.GetRequestID(ctx)

	var verr *ValidationError
	if errors.As(err, &verr) {
		status = fiber.StatusBadRequest
		gocolorlog.Warnf(
			"[HTTP %d] Validation error: %v | %s %s | IP: %s | RequestID: %s",
			status, verr.Msg, method, path, ip, reqID,
		)
		return ctx.Status(status).JSON(base.NewErrorResponse(status, verr.Msg))
	}

	var fiberErr *fiber.Error
	if errors.As(err, &fiberErr) {
		status = fiberErr.Code
		gocolorlog.Warnf(
			"[HTTP %d] Fiber error: %v | %s %s | IP: %s | RequestID: %s",
			status, fiberErr.Message, method, path, ip, reqID,
		)
		return ctx.Status(status).JSON(base.NewErrorResponse(status, fiberErr.Message))
	}

	msg := "Internal Server Error"
	if showDetails {
		msg = err.Error()
	}
	gocolorlog.Errorf(
		"[HTTP %d] Internal error: %v | %s %s | IP: %s | RequestID: %s",
		status, msg, method, path, ip, reqID,
	)
	return ctx.Status(status).JSON(base.NewErrorResponse(status, msg))
}
