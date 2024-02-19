package repositories

import (
	"context"
	"github.com/nunoonu/file-upload-service/internal/core/ports"
	"log/slog"
	"os"
	"path/filepath"
)

type fileRepository struct {
	Path string
}

func NewFileRepository(path string) ports.FileRepository {
	return &fileRepository{Path: path}
}

func (f fileRepository) Store(_ context.Context, fileName string, file []byte) error {

	path := filepath.Join(f.Path, fileName)
	newFile, err := os.Create(path)
	if err != nil {
		return err
	}
	defer func(newFile *os.File) {
		err = newFile.Close()
		if err != nil {
			slog.Error("Failed closing file", slog.String("Err", err.Error()))
			return
		}
	}(newFile)
	_, err = newFile.Write(file)

	return err
}
