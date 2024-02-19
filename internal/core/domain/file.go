package domain

import (
	"encoding/binary"
	"fmt"
)

const maxUploadSize = 2 * 1024 * 1024 // 2 mb

type UploadFileRequest struct {
	File     []byte
	FileName string
}

func (r UploadFileRequest) Validate() error {
	if binary.Size(r.File) > maxUploadSize {
		return fmt.Errorf("file size exceeds the limit of %d", maxUploadSize)
	}
	return nil
}

type UploadFileResponse struct {
}
