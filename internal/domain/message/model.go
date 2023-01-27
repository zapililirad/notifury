package message

import (
	"time"

	"github.com/zapililirad/notifury/internal/domain/access"
)

type Message struct {
	Text       string
	Class      *MessageClass
	RecievedAt time.Time
}

func NewMessage(class *MessageClass, text string) *Message {
	return &Message{
		Text:       text,
		Class:      class,
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
