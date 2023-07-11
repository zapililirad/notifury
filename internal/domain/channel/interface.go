package channel

import "github.com/zapililirad/notifury/internal/domain/message"

type NotificationChannel interface {
	Notify(message.Message) error
	GetUserUUID() string
	GetNotificationChannelType() NotificationChannelType
}
