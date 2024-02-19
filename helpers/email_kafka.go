package helpers

import (
	"context"
	"github.com/segmentio/kafka-go"
	"log/slog"
)

type MailParams struct {
	URL   string
	Topic string
}

func NewMailKafkaParams() *MailParams {
	return &MailParams{
		URL:   "localhost:9092",
		Topic: "email",
	}
}

func NewKafka(param *MailParams) *kafka.Conn {
	conn, err := kafka.DialLeader(context.Background(), "tcp", param.URL, param.Topic, 0)
	if err != nil {
		slog.Error("Failed to connect to kafka", slog.String("Err", err.Error()))
		panic(err.Error())
	}
	return conn
}
