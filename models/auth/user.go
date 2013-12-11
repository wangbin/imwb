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
	UserTable       = "auth_user"
	AnonymousUserId = "-1"
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
	if user.IsAnonymous() {
		return errors.New("Can't save anonymous user")
	}
	if _, ok := user.Validate(); !ok {
		return errors.New("Invalid user")
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

func (user *User) Validate() (map[string][]error, bool) {
	errMap := make(map[string][]error)
	valid := validation.Validation{}

	for _, v := range []*validation.ValidationResult{
		valid.Required(user.UserName, "username"),
		valid.MaxSize(user.UserName, 30, "usernameMax"),
		valid.Match(user.UserName, UserNamePattern, "usernamePattern"),
	} {
		if !v.Ok {
			if _, ok := errMap["username"]; ok {
				errMap["username"] = append(errMap["username"],
					errors.New(v.Error.Message))
			} else {
				errMap["username"] = []error{errors.New(v.Error.Message)}
			}
		}
	}
	if len(user.FirstName) > 0 {
		if v := valid.MaxSize(user.FirstName, 30, "first name"); !v.Ok {
			errMap["firsname"] = []error{errors.New(v.Error.Message)}
		}
	}
	if len(user.LastName) > 0 {
		if v := valid.MaxSize(user.LastName, 30, "last name"); !v.Ok {
			errMap["lastname"] = []error{errors.New(v.Error.Message)}
		}
	}
	if len(user.Email) > 0 {
		if v := valid.Email(user.Email, "email"); !v.Ok {
			errMap["email"] = []error{errors.New(v.Error.Message)}
		}
	}
	return errMap, len(errMap) == 0
}

func (user *User) SetPassword(rawPass string) {
	user.Password = hashers.MakePassword(rawPass)
}

func (user *User) SetEmail(email string) {
	user.Email = NormalizeEmail(email)
}

func (user *User) IsAnonymous() bool {
	return user.Id == AnonymousUserId
}

func (user *User) CheckPassword(password string) bool {
	return hashers.CheckPassword(password, user.Password)
}

func GetUser(session *r.Session, userId string) *User {
	var user *User
	err := r.Table(UserTable).Get(userId).Run(session).One(&user)
	if err != nil {
		return NewAnonymousUser()
	}
	if user == nil {
		user = NewAnonymousUser()
	}
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
