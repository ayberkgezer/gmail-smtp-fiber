package middleware

import (
	"time"

	"github.com/ayberkgezer/gocolorlog"
	"github.com/gofiber/fiber/v2"
)

func Logger(ctx *fiber.Ctx) error {
	start := time.Now()
	err := ctx.Next()
	latency := time.Since(start)
	status := ctx.Response().StatusCode()
	method := ctx.Method()
	path := ctx.OriginalURL()
	ip := ctx.IP()
	reqID := GetRequestID(ctx)

	if err == nil {
		gocolorlog.HTTP(status, method, path, latency, ip, reqID, nil)
	}

	return err
}
