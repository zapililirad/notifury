package message

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zapililirad/notifury/internal/domain/access"
)

func TestMessageClass_GetName(t *testing.T) {
	type fields struct {
		UUID      string
		ClassName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get back message class name",
			fields: fields{
				UUID:      uuid.NewString(),
				ClassName: "Test message class",
			},
			want: "Test message class",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &MessageClass{
				UUID:      tt.fields.UUID,
				ClassName: tt.fields.ClassName,
			}
			assert.Equal(t, tt.want, c.GetName())
		})
	}
}

func TestMessageClass_GetSecurityUUID(t *testing.T) {
	type fields struct {
		UUID      string
		ClassName string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get back message class uuid",
			fields: fields{
				UUID:      uuid.NewString(),
				ClassName: "Test message class",
			},
			want: "valid uuid string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &MessageClass{
				UUID:      tt.fields.UUID,
				ClassName: tt.fields.ClassName,
			}
			assert.Equal(t, c.UUID, c.GetSecurityUUID())
			assert.NotPanics(t, func() { uuid.MustParse(c.GetSecurityUUID()) })
		})
	}
}

func TestMessageClass_GetSecurityPrincipalType(t *testing.T) {
	type fields struct {
		UUID      string
		ClassName string
	}
	tests := []struct {
		name   string
		fields fields
		want   access.SecurityPrincipalType
	}{
		{
			name: "Get back message class uuid",
			fields: fields{
				UUID:      uuid.NewString(),
				ClassName: "Test message class",
			},
			want: access.MessageClass,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &MessageClass{
				UUID:      tt.fields.UUID,
				ClassName: tt.fields.ClassName,
			}
			assert.Equal(t, tt.want, c.GetSecurityPrincipalType())
		})
	}
}
