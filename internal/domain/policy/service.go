package policy

import (
	"context"

	"github.com/zapililirad/notifury/internal/app"
	"github.com/zapililirad/notifury/internal/domain/message"
)

type DistributionPolicyService struct {
	storage Storage
}

func (s *DistributionPolicyService) GetDistributionPoliciesByMessageClass(ctx context.Context, class *message.MessageClass) (*DistributionPolicy, error) {
	return nil, app.ErrNotImplemented
}

func (s *DistributionPolicyService) GetAllDistributionPolicies(ctx context.Context, limit, offset int) []*DistributionPolicy {
	return nil
}
