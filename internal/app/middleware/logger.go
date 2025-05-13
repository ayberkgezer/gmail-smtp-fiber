package middleware

import (
	"log"
	"time"

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

	var color, level string
	switch {
	case status >= 500:
		color = colorRed
		level = "ERROR"
	case status >= 400:
		color = colorYellow
		level = "WARN"
	default:
		color = colorGreen
		level = "INFO"
	}

	log.Printf(
		"%s[%s]%s %s %s → %d (%s)",
		color, level, colorReset,
		method, path, status, latency,
	)

	if err != nil {
		log.Printf(
			"%s[ERROR]%s handler error: %v",
			colorRed, colorReset, err,
		)
	}
	return err
}
