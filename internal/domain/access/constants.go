package access

type SecurityPrincipalType int

const (
	User SecurityPrincipalType = iota
	Group
	MessageClass
	DistributionPolicy
)
