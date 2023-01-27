package channel

import (
	"fmt"

	"github.com/zapililirad/notifury/internal/domain/message"
	"github.com/zapililirad/notifury/internal/domain/user"
)

type EmailChannel struct {
	User *user.User
}

func (c *EmailChannel) PushMessage(message message.Message) error {
	// TODO: Implement email notications

	_, err := fmt.Printf("Email for user %s: %s", c.User.GetName(), message.Text)

	return err
}
