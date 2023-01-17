package user

import (
	"time"

	"github.com/zapililirad/notifury/internal/domain/access"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UUID      string
	FirstName string
	LastName  string
	Email     string
	Active    bool
	password  password
}

func (u *User) GetName() string {
	return u.FirstName + " " + u.LastName
}

func (u *User) GetSecurityUUID() string {
	return u.UUID
}

func (u *User) GetSecurityPrincipalType() access.SecurityPrincipalType {
	return access.User
}

type password struct {
	password string
	created  time.Time
}

func SetPassword(s string) password {
	//TODO: Implement password policy check
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return password{} //TODO: Check if this is secure solution
	}

	return password{
		created:  time.Now(),
		password: string(hashedPassword),
	}
}

func (p *password) CompareWithString(s string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.password), []byte(s))
}
