package auth

import (
	"fmt"
	r "github.com/christopherhesse/rethinkgo"
	"strings"
	"time"
)

func NormalizeEmail(email string) string {
	if len(email) == 0 {
		return email
	}
	index := strings.LastIndex(email, "@")
	if index == -1 {
		return email
	}
	prefix, suffix := email[:index], email[index+1:]
	return fmt.Sprintf("%s@%s", prefix, strings.ToLower(suffix))
}

func NewUser(username string) *User {
	user := &User{UserName: username}
	user.DateJoined = time.Now()
	user.LastLogin = time.Now()
	return user
}

func NewAnonymousUser() *User {
	user := NewUser("AnonymousUser")
	user.Id = AnonymousUserId
	return user
}

func Authenticate(session *r.Session, name, password string) (*User, bool) {
	var users []*User
	err := r.Table(UserTable).GetAll("username", name).Run(session).All(&users)
	if err != nil {
		return nil, false
	}
	if len(users) == 0 {
		return nil, false
	}
	user := users[0]
	if !user.CheckPassword(password) {
		return nil, false
	}
	return user, true
}
