package group

import "context"

type Storage interface {
	SaveGroup(ctx context.Context, g *Group) error
	// GetGroup(ctx context.Context, uuid string) (*Group, error)
	// GetAllGroups(ctx context.Context, limit, offset int) []*Group
}
