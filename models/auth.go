package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/validation"
	r "github.com/dancannon/gorethink"
	"github.com/wangbin/imwb/utils/hashers"
	"regexp"
	"strings"
	"time"
)

const (
	UserTable            = "auth_user"
	AnonymousUserId      = "-1"
	UserTableSecondIndex = "username"
)

var (
	UserNamePattern = regexp.MustCompile(`^[[:word:]@+-]+$`)
)

type User struct {
	Id          string    `gorethink:"id,omitempty"`
	UserName    string    `gorethink:"username"`
	Password    string    `gorethink:"password"`
	FirstName   string    `gorethink:"first_name"`
	LastName    string    `gorethink:"last_name"`
	Email       string    `gorethink:"email"`
	IsStaff     bool      `gorethink:"is_staff"`
	IsSuperUser bool      `gorethink:"is_superuser"`
	IsActive    bool      `gorethink:"is_active"`
	DateJoined  time.Time `gorethink:"date_joined"`
	LastLogin   time.Time `gorethink:"last_login"`
	Groups      []string  `gorethink:"groups"`
}

func (user *User) mapping() map[string]interface{} {
	return map[string]interface{}{
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

	if len(user.Id) > 0 {
		_, err = r.Table(UserTable).Get(user.Id).Update(user.mapping()).RunWrite(session)
	} else {
		var response r.WriteResponse
		response, err = r.Table(UserTable).Insert(user).RunWrite(session)
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

func (user *User) IsAuthenticated() bool {
	return !user.IsAnonymous()
}

func (user *User) CheckPassword(password string) bool {
	return hashers.CheckPassword(password, user.Password)
}

func GetUser(session *r.Session, userId string) *User {
	var user *User
	row, err := r.Table(UserTable).Get(userId).RunRow(session)
	if err != nil || row.IsNil() {
		return NewAnonymousUser()
	}
	err = row.Scan(&user)
	if err != nil {
		return NewAnonymousUser()
	}
	return user
}

func Authenticate(session *r.Session, name, password string) (*User, bool) {
	var user *User
	rows, err := r.Table(UserTable).GetAllByIndex("username", name).RunRow(session)
	if err != nil || rows.IsNil() {
		fmt.Println(rows.IsNil())
		return nil, false
	}
	err = rows.Scan(&user)
	if err != nil || user == nil {
		return nil, false
	}
	if !user.CheckPassword(password) {
		return nil, false
	}
	return user, true
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
