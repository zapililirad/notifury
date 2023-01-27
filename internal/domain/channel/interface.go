package channel

import "github.com/zapililirad/notifury/internal/domain/message"

type NotificationChannel interface {
	PushMessage(message.Message) error
}
