package app

import (
	"github.com/nunoonu/file-upload-service/helpers"
	"github.com/nunoonu/file-upload-service/internal/core/usecases"
	"github.com/nunoonu/file-upload-service/internal/handlers"
	"github.com/nunoonu/file-upload-service/internal/repositories"
	"log/slog"
	"os"
)

func InitializeApp() *App {

	setLogLevel()

	params := helpers.NewMailKafkaParams()
	kCon := helpers.NewKafka(params)
	mailRepo := repositories.NewMailRepository(kCon)

	fileRepo := repositories.NewFileRepository("resource")
	fileUsc := usecases.NewFileUseCase(fileRepo, mailRepo)
	fileHdl := handlers.NewFileHandler(fileUsc)
	router := handlers.NewRouter(fileHdl)

	httpServParams := handlers.NewHTTPServiceParams()
	httpServ := handlers.NewHTTPService(httpServParams, router)
	return ProvideApp(httpServ)
}

func setLogLevel() {
	l := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(l)
}
