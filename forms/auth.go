package forms

import (
	"errors"
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/models/auth"
)

type LoginForm struct {
	Name           string `form:"username"`
	Password       string `form:"password,password,"`
	errorMap       map[string][]error
	nonFieldErrors []error
	rs             *r.Session
	user           *auth.User
}

func (this *LoginForm) IsValid() bool {
	var ok bool
	user := auth.NewUser(this.Name)
	this.errorMap, ok = user.Validate()
	if ok {
		this.user, ok = auth.Authenticate(this.rs, this.Name, this.Password)
		if !ok {
			this.errorMap["username"] = []error{
				errors.New("Please enter a correct name and password")}
		}
	}
	return ok
}

func (this *LoginForm) ErrorMap() map[string][]error {
	return this.errorMap
}

func (this *LoginForm) SetRs(rs *r.Session) {
	this.rs = rs
}

func (this *LoginForm) SetNonFieldError(err error) {
	this.nonFieldErrors = append(this.nonFieldErrors, err)
}

func (this *LoginForm) NonFieldErrors() []error {
	return this.nonFieldErrors
}

func (this *LoginForm) User() *auth.User {
	return this.user
}

func NewLoginForm() *LoginForm {
	form := &LoginForm{nonFieldErrors: make([]error, 0),
		errorMap: make(map[string][]error)}
	return form
}
