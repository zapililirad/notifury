package access

type SecurityPrincipal interface {
	GetSecurityUUID() string
	GetSecurityPrincipalType() SecurityPrincipalType
}
