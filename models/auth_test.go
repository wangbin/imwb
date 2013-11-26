package models

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
	u.UserName = "admin"
	if err := u.Validate(); err != nil {
		t.Error(err)
	}
	u.UserName = "admin#"
	if err:=u.Validate();err == nil {
		t.Error("username contains invalid character")
	}
}
