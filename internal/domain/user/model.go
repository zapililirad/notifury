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
	Email     string //Unique, login
	IsActive  bool
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

func setPassword(p *password, s string) error {
	//TODO: Implement password policy check
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		return err //TODO: Check if this is secure solution
	}

	p.password = string(hashedPassword)
	p.created = time.Now()

	return nil
}

func (p *password) CompareWithString(s string) error {
	return bcrypt.CompareHashAndPassword([]byte(p.password), []byte(s))
}
