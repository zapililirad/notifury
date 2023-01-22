package group

import "github.com/zapililirad/notifury/internal/domain/access"

type Group struct {
	UUID      string
	GroupName string
	Members   []access.SecurityPrincipal
}

func (g *Group) GetName() string {
	return g.GroupName
}

func (g *Group) GetSecurityUUID() string {
	return g.UUID
}

func (g *Group) GetSecurityPrincipalType() access.SecurityPrincipalType {
	return access.Group
}

func (g Group) IsEmpty() bool {
	return len(g.Members) == 0
}
