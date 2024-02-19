package repositories

import (
	"context"
	"encoding/json"
	"github.com/nunoonu/file-upload-service/internal/core/ports"
	"github.com/segmentio/kafka-go"
)

type Mail struct {
	FileName string
	File     []byte
}

type mailRepository struct {
	kafkaCon *kafka.Conn
}

func NewMailRepository(kCon *kafka.Conn) ports.MailRepository {
	return &mailRepository{kafkaCon: kCon}
}

func (m mailRepository) Send(_ context.Context, fileName string, file []byte) error {
	ma := Mail{File: file, FileName: fileName}
	km := kafka.Message{Value: compressToJsonBytes(ma)}
	_, err := m.kafkaCon.WriteMessages(km)
	return err
}

func compressToJsonBytes(obj any) []byte {
	raw, _ := json.Marshal(obj)
	return raw
}
