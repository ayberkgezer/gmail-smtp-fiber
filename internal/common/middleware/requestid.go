package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

const RequestIDHeader = "X-Request-ID"

func RequestID(ctx *fiber.Ctx) error {
	reqID := ctx.Get(RequestIDHeader)
	if reqID == "" {
		reqID = uuid.New().String()
	}
	ctx.Set(RequestIDHeader, reqID)
	ctx.Locals(RequestIDHeader, reqID)
	return ctx.Next()
}

func GetRequestID(ctx *fiber.Ctx) string {
	if v := ctx.Locals(RequestIDHeader); v != nil {
		return v.(string)
	}
	return ""
}
