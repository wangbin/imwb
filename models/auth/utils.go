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

func NewUser(username string) (*User, error) {
	user := &User{UserName: username}
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	user.DateJoined = time.Now()
	user.LastLogin = time.Now()
	return user, nil
}
