package auth

import (
	"errors"
	"github.com/astaxie/beego/validation"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/utils/hashers"
	"regexp"
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

func (user *User) SetEmail(email string) {
	user.Email = NormalizeEmail(email)
}
