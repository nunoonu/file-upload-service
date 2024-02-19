package ports

import (
	"context"
	"github.com/nunoonu/file-upload-service/internal/core/domain"
)

type FileUseCase interface {
	Upload(ctx context.Context, req *domain.UploadFileRequest) (*domain.UploadFileResponse, error)
}
