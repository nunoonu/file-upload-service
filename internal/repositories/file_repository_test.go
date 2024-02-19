package repositories

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FileRepositoryTest struct {
	suite.Suite
}

func (suite *FileRepositoryTest) SetupTest() {
}

func TestFileRepository(t *testing.T) {
	suite.Run(t, new(FileRepositoryTest))
}

func (suite *FileRepositoryTest) TestSuccessPersistFile() {
	suite.Run("Success", func() {
		repository := NewFileRepository("/tmp")
		err := repository.Store(context.TODO(), "test.txt", []byte("test"))
		assert.NoError(suite.T(), err)
	})
}
