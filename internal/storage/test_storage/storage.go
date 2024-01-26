package storage

import (
	"context"

	"github.com/zapililirad/notifury/internal/app"
	"github.com/zapililirad/notifury/internal/domain/group"
	// "github.com/zapililirad/notifury/internal/domain/group"
)

type TestStorage struct {
	groups map[string]*struct {
		name    string
		members []string
	}
}

func New() *TestStorage {
	return &TestStorage{}
}

func (s *TestStorage) SaveGroup(ctx context.Context, g *group.Group) error {
	s.groups[g.UUID] = &struct {
		name    string
		members []string
	}{
		name:    g.GroupName,
		members: nil,
	}

	for _, sp := range g.Members {
		s.groups[g.UUID].members = append(s.groups[g.UUID].members, sp.GetSecurityUUID())
	}

	return app.ErrNotImplemented
}
