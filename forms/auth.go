package forms

import (
	"github.com/wangbin/imwb/models/auth"
)

type LoginForm struct {
	Name           string `form:"username"`
	Password       string `form:"password,password,"`
	errorMap       map[string][]error
	nonFieldErrors []error
}

func (this *LoginForm) IsValid() bool {
	var ok bool
	this.errorMap, ok = this.validate()
	return ok
}

func (this *LoginForm) validate() (map[string][]error, bool) {
	user := auth.NewUser(this.Name)
	return user.Validate()
}

func (this *LoginForm) ErrorMap() map[string][]error {
	return this.errorMap
}

func (this *LoginForm) SetNonFieldError(err error) {
	this.nonFieldErrors = append(this.nonFieldErrors, err)
}

func (this *LoginForm) NonFieldErrors() []error {
	return this.nonFieldErrors
}

func NewLoginForm() *LoginForm {
	form := &LoginForm{nonFieldErrors: make([]error, 0),
		errorMap: make(map[string][]error)}
	return form
}
