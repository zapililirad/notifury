package group

import (
	"context"

	"github.com/google/uuid"
	"github.com/zapililirad/notifury/internal/app"
	"github.com/zapililirad/notifury/internal/domain/access"
	"github.com/zapililirad/notifury/internal/domain/user"
)

// type Service interface {
// 	GetGroupByUUID(ctx context.Context, uuid string) *Group
// 	GetAllGroups(ctx context.Context, limit, offset int) []*Group
// 	AppendWith(ctx context.Context, group *Group, sp access.SecurityPrincipal) error
// }

type GroupService struct {
	storage Storage
}

func (s *GroupService) GetGroupByUUID(ctx context.Context, uuid string) (*Group, error) {
	return nil, app.ErrNotImplemented
}

func (s *GroupService) GetAllGroups(ctx context.Context, limit, offset int) []*Group {
	return nil
}

func (s *GroupService) AppendToGroup(ctx context.Context, g *Group, sp access.SecurityPrincipal) error {
	if g.GetSecurityUUID() == sp.GetSecurityUUID() {
		return ErrAppendSelf
	}

	if sp.GetSecurityPrincipalType() == access.Group && s.HasGroupMemberRecursive(ctx, sp.(*Group), g) {
		return ErrRecursiveNesting
	}

	g.Members = append(g.Members, sp)
	return nil
}

func (s *GroupService) RemoveFromGroup(ctx context.Context, g *Group, sp access.SecurityPrincipal) error {
	for i, m := range g.Members {
		if m.GetSecurityUUID() == sp.GetSecurityUUID() {
			g.Members = append(g.Members[:i], g.Members[i+1:]...)
			return nil
		}
	}

	return ErrNotContained
}

func (s *GroupService) GetAllGroupMembers(ctx context.Context, g *Group) []access.SecurityPrincipal {
	return g.Members
}

func (s *GroupService) HasGroupMember(ctx context.Context, g *Group, sp access.SecurityPrincipal) bool {
	// TODO: Maybe delete?

	for _, m := range g.Members {
		if m.GetSecurityUUID() == sp.GetSecurityUUID() {
			return true
		}
	}

	return false
}

func (s *GroupService) HasGroupMemberRecursive(ctx context.Context, g *Group, sp access.SecurityPrincipal) bool {
	for _, m := range g.Members {
		if m.GetSecurityUUID() == sp.GetSecurityUUID() {
			return true
		}

		if m.GetSecurityPrincipalType() == access.Group {
			if s.HasGroupMemberRecursive(ctx, m.(*Group), sp) {
				return true
			}
		}
	}

	return false
}

func (s *GroupService) GetAllUsersRecursive(ctx context.Context, g *Group) []*user.User {
	users := []*user.User{}

	for _, m := range g.Members {
		switch m.GetSecurityPrincipalType() {
		case access.User:
			users = append(users, m.(*user.User))
		case access.Group:
			users = append(users, s.GetAllUsersRecursive(ctx, m.(*Group))...)
		}
	}

	return users
}

func (s *GroupService) CreateEmptyGroup(ctx context.Context, name string) (*Group, error) {
	if err := app.ValidateName(name); err != nil {
		return nil, err
	}

	g := &Group{
		UUID:      uuid.NewString(),
		GroupName: name,
		Members:   nil,
	}

	if err := s.storage.SaveGroup(ctx, g); err != nil {
		return nil, err
	}

	return g, nil
}

func (s *GroupService) DeleteGroup(ctx context.Context, g *Group) error {
	return app.ErrNotImplemented
}
