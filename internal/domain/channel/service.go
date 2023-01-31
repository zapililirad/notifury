package channel

import (
	"context"

	"github.com/zapililirad/notifury/internal/domain/user"
)

type NotificationChannelService struct {
	storage Storage
}

func (s *NotificationChannelService) GetNotificationChannelsByUser(ctx context.Context, u *user.User) []NotificationChannel {
	// TODO: Implement GetNotificationChannelsByUser
	return nil
}
