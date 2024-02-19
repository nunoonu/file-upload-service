package app

import (
	"github.com/nunoonu/file-upload-service/internal/handlers"
)

func ProvideApp(httpServer *handlers.HTTPService) *App {
	return &App{
		HTTPServer: *httpServer,
	}
}

type App struct {
	HTTPServer handlers.HTTPService
}
