package ports

import "context"

type FileRepository interface {
	Store(ctx context.Context, fileName string, file []byte) error
}

type MailRepository interface {
	Send(ctx context.Context, fileName string, file []byte) error
}
