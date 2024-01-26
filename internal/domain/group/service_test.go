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

// func TestGroupService_CreateEmptyGroup(t *testing.T) {
// 	storage := storage.New()
// 	type args struct {
// 		ctx  context.Context
// 		name string
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *Group
// 		wantErr error
// 	}{
// 		{
// 			name:    "Create group",
// 			args:    args{ctx: context.Background(), name: "Test group"},
// 			want:    &Group{Members: nil, GroupName: "Test group", UUID: uuid.NewString()},
// 			wantErr: nil,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &GroupService{
// 				storage: storage,
// 			}
// 			got, err := s.CreateEmptyGroup(tt.args.ctx, tt.args.name)

// 			if tt.wantErr != nil {
// 				assert.EqualError(t, err, "")
// 			} else {
// 				assert.NoError(t, err)
// 			}

// 			if tt.want != nil {
// 				assert.Len(t, got.Members, 0)
// 				assert.Equal(t, tt.want.GroupName, got.GroupName)
// 				assert.NotPanics(t, func() { uuid.MustParse(got.UUID) }, "Group id is not UUID")
// 			} else {
// 				assert.Nil(t, got)
// 			}
// 		})
// 	}
// }

// func TestGroupService_GetAllUsersRecursive(t *testing.T) {
// 	type fields struct {
// 		storage Storage
// 	}
// 	type args struct {
// 		ctx context.Context
// 		g   *Group
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   []*user.User
// 	}{
// 		{
// 			name: "Empty group",
// 			args: args{ctx: context.Background(), g: TestGroup(t)},
// 			want: []*user.User{},
// 		},
// 		{
// 			name: "Group with one user",
// 			args: args{ctx: context.Background(), g: &Group{
// 				GroupName: "Test group",
// 				Members:   []access.SecurityPrincipal{user.TestUser(t)},
// 			}},
// 			want: []*user.User{user.TestUser(t)},
// 		},
// 		{
// 			name: "Nested group with one user",
// 			args: args{ctx: context.Background(), g: &Group{
// 				GroupName: "Test group",
// 				Members: []access.SecurityPrincipal{&Group{
// 					GroupName: "Test group 2",
// 					Members:   []access.SecurityPrincipal{user.TestUser(t)},
// 				}},
// 			}},
// 			want: []*user.User{user.TestUser(t)},
// 		},
// 		{
// 			name: "Two groups nested, with users",
// 			args: args{ctx: context.Background(), g: &Group{
// 				GroupName: "Test group",
// 				Members: []access.SecurityPrincipal{
// 					&Group{
// 						GroupName: "Test group 2",
// 						Members:   []access.SecurityPrincipal{user.TestUser(t)},
// 					},
// 					&Group{
// 						GroupName: "Test group 3",
// 						Members:   []access.SecurityPrincipal{user.TestUser(t)},
// 					}},
// 			}},
// 			want: []*user.User{user.TestUser(t), user.TestUser(t)},
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			s := &GroupService{
// 				storage: tt.fields.storage,
// 			}
// 			if got := s.GetAllUsersRecursive(tt.args.ctx, tt.args.g); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("GroupService.GetAllUsersRecursive() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
