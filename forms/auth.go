package forms

import (
	"errors"
	"github.com/wangbin/imwb/models"
)

type LoginForm struct {
	Name           string `form:"username"`
	Password       string `form:"password,password,"`
	errorMap       map[string][]error
	nonFieldErrors []error
	user           *models.User
}

func (this *LoginForm) IsValid() bool {
	var ok bool
	user := models.NewUser(this.Name)
	this.errorMap, ok = user.Validate()
	if ok {
		this.user, ok = models.Authenticate(this.Name, this.Password)
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

func (this *LoginForm) SetNonFieldError(err error) {
	this.nonFieldErrors = append(this.nonFieldErrors, err)
}

func (this *LoginForm) NonFieldErrors() []error {
	return this.nonFieldErrors
}

func (this *LoginForm) User() *models.User {
	return this.user
}

func NewLoginForm(user *models.User) *LoginForm {
	form := &LoginForm{nonFieldErrors: make([]error, 0),
		errorMap: make(map[string][]error), user: user}
	return form
}
