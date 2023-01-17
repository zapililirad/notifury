package message

import (
	"time"

	"github.com/google/uuid"
	"github.com/zapililirad/notifury/internal/domain/access"
)

type Message struct {
	Text       string
	Class      MessageClass
	RecievedAt time.Time
}

func NewMessage(text string) Message {
	return Message{
		Text: text,
		// TODO: Add constructor for MessageClass
		Class: MessageClass{
			UUID:      uuid.NewString(),
			ClassName: "Generic",
		},
		RecievedAt: time.Now(),
	}
}

type MessageClass struct {
	UUID      string
	ClassName string
	// TODO: Add more parameters to describe classes
}

func (c *MessageClass) GetName() string {
	return c.ClassName
}

func (c *MessageClass) GetSecurityUUID() string {
	return c.UUID
}

func (c *MessageClass) GetSecurityPrincipalType() access.SecurityPrincipalType {
	return access.MessageClass
}
