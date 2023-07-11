package channel

import (
	"fmt"

	"github.com/zapililirad/notifury/internal/domain/message"
	"github.com/zapililirad/notifury/internal/domain/user"
)

type EmailChannel struct {
	User     *user.User
	IsActive bool
	Email    string
}

func (c *EmailChannel) Notify(message message.Message) error {
	// TODO: Implement email notications

	_, err := fmt.Printf("Email for user %s: %s", c.User.GetName(), message.Text)

	return err
}

func (c *EmailChannel) GetNotificationChannelType() NotificationChannelType {
	return Email
}

func (c *EmailChannel) GetUserUUID() string {
	return c.User.GetSecurityUUID()
}
