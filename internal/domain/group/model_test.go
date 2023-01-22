package group

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zapililirad/notifury/internal/domain/access"
)

func TestGroup_GetName(t *testing.T) {
	g := TestGroup(t)
	assert.Equal(t, "Test group", g.GetName())
}

func TestGroup_GetSecurityUUID(t *testing.T) {
	g := TestGroup(t)
	assert.NotPanics(t, func() { uuid.MustParse(g.GetSecurityUUID()) })
}

func TestGroup_GetSecurityPrincipalType(t *testing.T) {
	g := TestGroup(t)
	assert.Equal(t, access.Group, g.GetSecurityPrincipalType())
}
