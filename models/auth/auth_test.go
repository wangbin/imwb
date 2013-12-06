package auth

import (
	"testing"
)

func TestUserValidate(t *testing.T) {
	u := &User{UserName: ""}
	if err := u.Validate(); err == nil {
		t.Error("username is empty")
	}
	u.UserName = "this is a long string that length is definitely longer than 30 characters"
	if err := u.Validate(); err == nil {
		t.Error("username max length should be 30")
	}
	u.UserName = "admin#"
	if err := u.Validate(); err == nil {
		t.Error("username contains invalid character")
	}
	u.UserName = "admin"
	if err := u.Validate(); err != nil {
		t.Error(err)
	}
	u.FirstName = "this is a long string that length is definitely longer than 30 characters"
	if err := u.Validate(); err == nil {
		t.Error("first name max length should be 30")
	}
	u.FirstName = "admin"
	u.Email = "admin@gmail"
	if err := u.Validate(); err == nil {
		t.Error("invalid email")
	}
	u.Email = "admin.system_it@gmail.com"
	if err := u.Validate(); err != nil {
		t.Error(err)
	}

}

func TestNormailzeEmail(t *testing.T) {
	email := ""
	if NormalizeEmail(email) != email {
		t.Error("empty email should return unchanged")
	}
	email = "admin#example.com"
	if NormalizeEmail(email) != email {
		t.Error("invalid email should return unchanged")
	}
	email = "admin@Example.com"
	emailNormalized := "admin@example.com"
	if NormalizeEmail(email) != emailNormalized {
		t.Error("email domain part should be lowercase")
	}

}
