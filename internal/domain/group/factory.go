package group

import (
	"github.com/google/uuid"
	"github.com/zapililirad/notifury/internal/domain/access"
)

type GroupFactory struct{}

func (f *GroupFactory) CreateEmptyGroup(name string) Group {
	return Group{
		GroupName:          name,
		UUID:               uuid.NewString(),
		SecurityPrincipals: []*access.SecurityPrincipal{},
	}
}
