package auth

import (
	"fmt"
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
