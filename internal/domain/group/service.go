package group

import (
	"context"

	"github.com/zapililirad/notifury/internal/domain/access"
)

type Service interface {
	GetGroupByUUID(ctx context.Context, uuid string) *Group
	GetAllGroups(ctx context.Context, limit, offset int) []*Group
	AppendWith(ctx context.Context, group *Group, sp access.SecurityPrincipal) error
}
