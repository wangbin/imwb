package auth

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/utils/hashers"
	"regexp"
	"strings"
	"time"
)

const (
	UserTable = "auth_user"
)

var (
	UserNamePattern = regexp.MustCompile(`^[[:word:]@+-]+$`)
)

type User struct {
	Id          string    `json:"id,omitempty"`
	UserName    string    `json:"username"`
	Password    string    `json:"password"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	Email       string    `json:"email"`
	IsStaff     bool      `json:"is_staff"`
	IsSuperUser bool      `json:"is_superuser"`
	IsActive    bool      `json:"is_active"`
	DateJoined  time.Time `json:"date_joined"`
	LastLogin   time.Time `json:"last_login"`
	Groups      []string  `json:"groups"`
}

func (user *User) mapping() r.Map {
	return r.Map{
		"username":     user.UserName,
		"password":     user.Password,
		"first_name":   user.FirstName,
		"last_name":    user.LastName,
		"email":        user.Email,
		"is_superuser": user.IsSuperUser,
		"is_staff":     user.IsStaff,
		"is_active":    user.IsActive,
		"date_joined":  user.DateJoined,
		"last_login":   user.LastLogin,
		"groups":       user.Groups,
	}
}

func (user *User) Save(session *r.Session) error {
	var err error
	if err = user.Validate(); err != nil {
		return err
	}

	var response r.WriteResponse
	if len(user.Id) > 0 {
		err = r.Table(UserTable).Get(user.Id).Update(user.mapping()).Run(session).One(
			&response)
	} else {
		err = r.Table(UserTable).Insert(user).Run(session).One(&response)
		if err == nil && len(response.GeneratedKeys) > 0 {
			user.Id = response.GeneratedKeys[0]
		}
	}
	return err
}

func (user *User) Validate() error {
	var v *validation.ValidationResult
	valid := validation.Validation{}
	if v = valid.Required(user.UserName, "username"); !v.Ok {
		return errors.New(v.Error.Message)
	}
	if v = valid.MaxSize(user.UserName, 30, "username"); !v.Ok {
		return errors.New(v.Error.Message)
	}
	if v = valid.Match(user.UserName, UserNamePattern, "username"); !v.Ok {
		return errors.New(v.Error.Message)
	}
	if len(user.FirstName) > 0 {
		if v = valid.MaxSize(user.FirstName, 30, "first name"); !v.Ok {
			return errors.New(v.Error.Message)
		}
	}
	if len(user.LastName) > 0 {
		if v = valid.MaxSize(user.LastName, 30, "last name"); !v.Ok {
			return errors.New(v.Error.Message)
		}
	}
	if len(user.Email) > 0 {
		if v = valid.Email(user.Email, "email"); !v.Ok {
			return errors.New(v.Error.Message)
		}
	}
	return nil
}

func (user *User) SetPassword(rawPass string) {
	user.Password = hashers.MakePassword(rawPass)
}

func CreateUser(username string) (*User, error) {
	user := &User{UserName: username}
	err := user.Validate()
	if err != nil {
		return nil, err
	}
	return user, nil
}

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