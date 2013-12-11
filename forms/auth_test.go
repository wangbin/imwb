package forms

import (
	"testing"
)

func TestLoginForm(t *testing.T) {
	form := &LoginForm{Name: ""}
	if form.IsValid() {
		t.FailNow()
	}
}
