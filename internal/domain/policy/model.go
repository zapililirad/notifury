package policy

import (
	"github.com/zapililirad/notifury/internal/domain/access"
	"github.com/zapililirad/notifury/internal/domain/message"
)

type DistributionPolicy struct {
	MessageClass *message.MessageClass
	Subscribers  []access.SecurityPrincipal
}

func (p *DistributionPolicy) GetSecurityUUID() string {
	return p.MessageClass.UUID
}

func (p *DistributionPolicy) GetSecurityPrincipalType() access.SecurityPrincipalType {
	return access.DistributionPolicy
}
