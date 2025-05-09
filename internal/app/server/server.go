package server

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ayberkgezer/gmail-smtp-fiber/internal/app/config"
	"github.com/gofiber/fiber/v2"
)

type server struct {
	app *fiber.App
}

func NewServer(app *fiber.App) *server {
	return &server{
		app: app,
	}
}

func (s *server) StartHttpServer() {
	go func() {
		gracefullShutdown(s.app)
	}()
	if err := s.app.Listen(fmt.Sprintf(":%s", config.Port)); err != nil {
		fmt.Println("Error Starting Server:", err)
		panic("can not start server")
	}

}

func gracefullShutdown(app *fiber.App) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shutting Down Server...")
	_, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		fmt.Println("Error shutting down server:", err)
	} else {
		fmt.Println("Server shut down gracefully")
	}
}
