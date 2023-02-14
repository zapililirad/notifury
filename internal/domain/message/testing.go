package message

import "testing"

func TestMessageClass(t *testing.T) *MessageClass {
	return &MessageClass{
		ClassName: "Test message class name",
		UUID:      "9d4f08d0-6930-4c51-85d3-2d82df9603b9",
	}
}
