package policy

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zapililirad/notifury/internal/domain/access"
	"github.com/zapililirad/notifury/internal/domain/message"
)

func TestDistributionPolicy_GetSecurityUUID(t *testing.T) {
	type fields struct {
		MessageClass *message.MessageClass
		Subscribers  []access.SecurityPrincipal
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get back Security uuid",
			fields: fields{
				MessageClass: message.TestMessageClass(t),
				Subscribers:  nil,
			},
			want: "9d4f08d0-6930-4c51-85d3-2d82df9603b9",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DistributionPolicy{
				MessageClass: tt.fields.MessageClass,
				Subscribers:  tt.fields.Subscribers,
			}
			assert.Equal(t, tt.want, p.GetSecurityUUID())
		})
	}
}

func TestDistributionPolicy_GetSecurityPrincipalType(t *testing.T) {
	type fields struct {
		MessageClass *message.MessageClass
		Subscribers  []access.SecurityPrincipal
	}
	tests := []struct {
		name   string
		fields fields
		want   access.SecurityPrincipalType
	}{
		{
			name: "Get back Distribution policy principal type",
			fields: fields{
				MessageClass: message.TestMessageClass(t),
				Subscribers:  nil,
			},
			want: access.DistributionPolicy,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &DistributionPolicy{
				MessageClass: tt.fields.MessageClass,
				Subscribers:  tt.fields.Subscribers,
			}
			assert.Equal(t, tt.want, p.GetSecurityPrincipalType())
		})
	}
}
