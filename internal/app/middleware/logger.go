package middleware

import (
	"time"

	"github.com/ayberkgezer/gocolorlog"
	"github.com/gofiber/fiber/v2"
)

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
)

func Logger(ctx *fiber.Ctx) error {
	start := time.Now()
	err := ctx.Next()
	latency := time.Since(start)
	status := ctx.Response().StatusCode()
	method := ctx.Method()
	path := ctx.OriginalURL()

	gocolorlog.HTTP(status, method, path, latency, err)
	return err
}
