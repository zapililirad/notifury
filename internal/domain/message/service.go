package message

import (
	"context"

	"github.com/google/uuid"
	"github.com/zapililirad/notifury/internal/app"
)

type MessageClassService struct {
	storage Storage
}

func (s *MessageClassService) GetMessageClassByUUID(ctx context.Context, uuid string) (*MessageClass, error) {
	return nil, app.ErrNotImplemented
}

func (s *MessageClassService) GetAllMessageClasses(ctx context.Context, limit, offset int) []*MessageClass {
	return nil
}

func (s *MessageClass) CreateMessageClass(ctx context.Context, name string) (*MessageClass, error) {
	if err := app.ValidateName(name); err != nil {
		return nil, err
	}

	return &MessageClass{
		ClassName: name,
		UUID:      uuid.NewString(),
	}, nil
}
