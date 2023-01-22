package group

import (
	"context"

	"github.com/zapililirad/notifury/internal/app"
	"github.com/zapililirad/notifury/internal/domain/access"
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
	// lite implementation
	// TODO: Implement append logic

	if g.GetSecurityUUID() == sp.GetSecurityUUID() {
		return ErrAppendSelf
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
