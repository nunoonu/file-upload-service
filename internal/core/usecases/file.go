package usecases

import (
	"context"
	"github.com/nunoonu/file-upload-service/internal/core/domain"
	"github.com/nunoonu/file-upload-service/internal/core/ports"
	"log/slog"
)

type fileUseCase struct {
	fileRepository ports.FileRepository
	mailRepository ports.MailRepository
}

func NewFileUseCase(fileRepository ports.FileRepository, mailRepository ports.MailRepository) ports.FileUseCase {
	return &fileUseCase{
		fileRepository: fileRepository,
		mailRepository: mailRepository}
}

func (r fileUseCase) Upload(ctx context.Context, req *domain.UploadFileRequest) (*domain.UploadFileResponse, error) {
	if err := req.Validate(); err != nil {
		slog.Error("Validation fail", slog.String("Err", err.Error()))
		return nil, err
	}
	if err := r.fileRepository.Store(ctx, req.FileName, req.File); err != nil {
		slog.Error("File store fail", slog.String("Err", err.Error()))
		return nil, err
	}
	if err := r.mailRepository.Send(ctx, req.FileName, req.File); err != nil {
		slog.Error("Mail send fail", slog.String("Err", err.Error()))
		return nil, err
	}
	return &domain.UploadFileResponse{}, nil
}
