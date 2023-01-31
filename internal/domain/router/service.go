package router

import (
	"context"

	"github.com/zapililirad/notifury/internal/domain/access"
	"github.com/zapililirad/notifury/internal/domain/channel"
	"github.com/zapililirad/notifury/internal/domain/group"
	"github.com/zapililirad/notifury/internal/domain/message"
	"github.com/zapililirad/notifury/internal/domain/policy"
	"github.com/zapililirad/notifury/internal/domain/user"
)

type RouterService struct {
	policyService  *policy.DistributionPolicyService
	groupService   *group.GroupService
	channelService *channel.NotificationChannelService
}

func NewRouterService(
	p *policy.DistributionPolicyService,
	g *group.GroupService,
	c *channel.NotificationChannelService,
) *RouterService {
	return &RouterService{
		policyService:  p,
		groupService:   g,
		channelService: c,
	}
}

func (s *RouterService) PushMessage(ctx context.Context, message *message.Message) error {
	// TODO: Make me pretty - need to refactor

	policy, err := s.policyService.GetDistributionPoliciesByMessageClass(ctx, message.Class)
	if err != nil {
		return err
	}

	users := []*user.User{}

	for _, sp := range policy.Subscribers {

		if sp.GetSecurityPrincipalType() == access.User {
			users = append(users, sp.(*user.User))
		}

		if sp.GetSecurityPrincipalType() == access.Group {
			users = append(users, s.groupService.GetAllUsersRecursive(ctx, sp.(*group.Group))...)
		}
	}

	for _, u := range users {
		channels := s.channelService.GetNotificationChannelsByUser(ctx, u)

		for _, c := range channels {
			c.PushMessage(*message)
		}

	}

	return nil
}
