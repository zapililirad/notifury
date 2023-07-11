package channel

import (
	"context"

	"github.com/zapililirad/notifury/internal/app"
	"github.com/zapililirad/notifury/internal/domain/user"
)

type NotificationChannelService struct {
	storage Storage
}

func (s *NotificationChannelService) GetNotificationChannelsByUser(ctx context.Context, u *user.User) []NotificationChannel {
	// TODO: Implement GetNotificationChannelsByUser
	return nil
}

func (s *NotificationChannelService) CreateNotificationChannel(ctx context.Context, u *user.User, channelType NotificationChannelType) (NotificationChannel, error) {
	switch channelType {
	case Email:
		return s.NewEmailChannel(ctx, u)
	default:
		return nil, app.ErrNotImplemented
	}
}

func (s *NotificationChannelService) NewEmailChannel(ctx context.Context, u *user.User) (*EmailChannel, error) {
	return &EmailChannel{
		User:     u,
		IsActive: true,
		Email:    u.Email,
	}, nil
}

func (s *NotificationChannelService) DeleteNotificationChannel(ctx context.Context, u *user.User, channelType NotificationChannelType) error {
	return app.ErrNotImplemented
}
