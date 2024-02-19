package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/nunoonu/file-upload-service/internal/core/domain"
	"github.com/nunoonu/file-upload-service/internal/core/ports"
	"io"
	"log/slog"
	"mime/multipart"
	"net/http"
)

type FileHandler struct {
	auc ports.FileUseCase
}

func NewFileHandler(auc ports.FileUseCase) *FileHandler {
	return &FileHandler{auc: auc}
}

func (h *FileHandler) Upload(ctx *gin.Context) {

	form, err := ctx.MultipartForm()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	files := form.File["file"]

	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "No file found"})
		return
	}

	fileName, file, err := readFile(files[0])
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.auc.Upload(ctx, &domain.UploadFileRequest{File: file, FileName: fileName})
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	return
}

func readFile(file *multipart.FileHeader) (string, []byte, error) {
	openedFile, _ := file.Open()

	binaryFile, err := io.ReadAll(openedFile)

	if err != nil {
		return "", nil, err
	}

	defer func(openedFile multipart.File) {
		err = openedFile.Close()
		if err != nil {
			slog.Error("Failed closing file", slog.String("File", file.Filename))
		}
	}(openedFile)
	return file.Filename, binaryFile, nil
}
