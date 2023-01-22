package group

import "testing"

func TestGroup(t *testing.T) *Group {
	return &Group{
		GroupName: "Test group",
		UUID:      "c4bfd565-9976-4ed2-912a-0471c3443a4e",
		Members:   nil,
	}
}
