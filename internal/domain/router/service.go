package router

import (
	"context"

	"github.com/zapililirad/notifury/internal/app"
	"github.com/zapililirad/notifury/internal/domain/access"
	"github.com/zapililirad/notifury/internal/domain/group"
	"github.com/zapililirad/notifury/internal/domain/message"
	"github.com/zapililirad/notifury/internal/domain/policy"
)

type RouterService struct {
	policyService *policy.DistributionPolicyService
	groupService  *group.GroupService
}

func (s *RouterService) PushMessage(ctx context.Context, message *message.Message) error {
	// TODO: Make me pretty

	policy, err := s.policyService.GetDistributionPoliciesByMessageClass(ctx, message.Class)
	if err != nil {
		return err
	}

	for _, sp := range policy.Subscribers {

		if sp.GetSecurityPrincipalType() == access.User {

		}

		if sp.GetSecurityPrincipalType() == access.Group {
			continue
		}
	}

	return app.ErrNotImplemented
}
