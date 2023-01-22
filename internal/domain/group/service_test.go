package group

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zapililirad/notifury/internal/domain/access"
	"github.com/zapililirad/notifury/internal/domain/user"
)

func TestGroupService_AppendWith(t *testing.T) {
	s := &GroupService{}

	testCases := []struct {
		name string
		sp   access.SecurityPrincipal
		ok   bool
	}{
		{
			name: "Append test user",
			sp:   user.TestUser(t),
			ok:   true,
		},
		{
			name: "Append self",
			sp:   TestGroup(t),
			ok:   false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			g := TestGroup(t)

			if tc.ok {
				assert.NoError(t, s.AppendToGroup(ctx, g, tc.sp))
				assert.Len(t, g.Members, 1)
			} else {
				assert.Error(t, s.AppendToGroup(ctx, g, tc.sp))
				assert.Empty(t, g.Members)
			}
		})
	}
}

func TestGroupService_RemoveWith(t *testing.T) {
	s := &GroupService{}

	testCases := []struct {
		name       string
		group      *Group
		sp         access.SecurityPrincipal
		lenOfSlice int
		ok         bool
	}{
		{
			name:       "Remove test user from empty group",
			group:      TestGroup(t),
			sp:         user.TestUser(t),
			lenOfSlice: 0,
			ok:         false,
		},
		{
			name: "Remove the only user",
			group: &Group{
				GroupName: "Test group",
				UUID:      "c4bfd565-9976-4ed2-912a-0471c3443a4e",
				Members: []access.SecurityPrincipal{
					user.TestUser(t),
				},
			},
			sp:         user.TestUser(t),
			ok:         true,
			lenOfSlice: 0,
		},
		{
			name: "Remove user from 2 members group",
			group: &Group{
				GroupName: "Test group",
				UUID:      "c4bfd565-9976-4ed2-912a-0471c3443a4e",
				Members: []access.SecurityPrincipal{
					user.TestUser(t),
					&user.User{},
				},
			},
			sp:         user.TestUser(t),
			ok:         true,
			lenOfSlice: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			ctx := context.Background()

			err := s.RemoveFromGroup(ctx, tc.group, tc.sp)

			if tc.ok {
				assert.NoError(t, err)
			} else {
				assert.Error(t, err)
			}

			assert.Len(t, tc.group.Members, tc.lenOfSlice)
		})
	}
}
