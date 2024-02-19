package usecases

import (
	"context"
	"errors"
	"github.com/nunoonu/file-upload-service/internal/core/domain"
	"github.com/nunoonu/file-upload-service/internal/core/ports"
	"github.com/nunoonu/file-upload-service/internal/core/ports/mocks"
	"github.com/nunoonu/file-upload-service/internal/repositories"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type FileTestSuite struct {
	suite.Suite
	fileRepository     ports.FileRepository
	mockFileRepository *mocks.FileRepository

	mailRepository     ports.MailRepository
	mockMailRepository *mocks.MailRepository
}

func (suite *FileTestSuite) SetupTest() {
	suite.fileRepository = repositories.NewFileRepository("/tmp")
	suite.mockFileRepository = &mocks.FileRepository{}

	suite.mailRepository = repositories.NewMailRepository(nil)
	suite.mockMailRepository = &mocks.MailRepository{}
}

func TestUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(FileTestSuite))
}

func (suite *FileTestSuite) TestUploadSuite() {
	suite.Run("Success", func() {
		suite.mockFileRepository.On("Store", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		suite.mockMailRepository.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()

		usecase := NewFileUseCase(suite.mockFileRepository, suite.mockMailRepository)
		actual, err := usecase.Upload(context.TODO(), &domain.UploadFileRequest{})
		assert.NoError(suite.T(), err)
		assert.Equal(suite.T(), &domain.UploadFileResponse{}, actual)
	})

	suite.Run("ErrorBigFile", func() {

		b := make([]byte, 10000000000)
		usecase := NewFileUseCase(suite.mockFileRepository, suite.mockMailRepository)
		_, err := usecase.Upload(context.TODO(), &domain.UploadFileRequest{File: b})
		assert.Error(suite.T(), err)
	})

	suite.Run("ErrorNotStore", func() {
		suite.mockFileRepository.On("Store", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("NotStore")).Once()

		usecase := NewFileUseCase(suite.mockFileRepository, suite.mockMailRepository)
		actual, err := usecase.Upload(context.TODO(), &domain.UploadFileRequest{})
		assert.Error(suite.T(), err)
		assert.Nil(suite.T(), actual)
	})

	suite.Run("ErrorNotSend", func() {
		suite.mockFileRepository.On("Store", mock.Anything, mock.Anything, mock.Anything).Return(nil).Once()
		suite.mockMailRepository.On("Send", mock.Anything, mock.Anything, mock.Anything).Return(errors.New("NotSend")).Once()

		usecase := NewFileUseCase(suite.mockFileRepository, suite.mockMailRepository)
		actual, err := usecase.Upload(context.TODO(), &domain.UploadFileRequest{})
		assert.Error(suite.T(), err)
		assert.Nil(suite.T(), actual)
	})
}
