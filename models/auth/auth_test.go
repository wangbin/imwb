package auth

import (
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/settings"
	"testing"
)

func TestUserValidate(t *testing.T) {
	u := &User{UserName: ""}
	if _, ok := u.Validate(); ok {
		t.Error("username is empty")
	}
	u.UserName = "this is a long string that length is definitely longer than 30 characters"
	if _, ok := u.Validate(); ok {
		t.Error("username max length should be 30")
	}
	u.UserName = "admin#"
	if _, ok := u.Validate(); ok {
		t.Error("username contains invalid character")
	}
	u.UserName = "admin"
	if _, ok := u.Validate(); !ok {
		t.Error("username should be valid")
	}
	u.FirstName = "this is a long string that length is definitely longer than 30 characters"
	if _, ok := u.Validate(); ok {
		t.Error("first name max length should be 30")
	}
	u.FirstName = "admin"
	u.Email = "admin@gmail"
	if _, ok := u.Validate(); ok {
		t.Error("invalid email")
	}
	u.Email = "admin.system_it@gmail.com"
	if _, ok := u.Validate(); !ok {
		t.Error(u.Email)
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

func TestNewAnonmousUser(t *testing.T) {
	u := NewAnonymousUser()
	if u.Id != AnonymousUserId {
		t.Error("User is not anonymous")
	}
}

func TestAuthenticate(t *testing.T) {
	rs, _ := r.Connect(settings.DbUri, settings.DbName)
	defer rs.Close()
	if _, ok := Authenticate(rs, "wangbin", "1234"); !ok {
		t.FailNow()
	}
	if _, ok := Authenticate(rs, "wangbin", "11111"); ok {
		t.FailNow()
	}
}

func TestGetUser(t *testing.T) {
	rs, _ := r.Connect(settings.DbUri, settings.DbName)
	defer rs.Close()
	user := GetUser(rs, "somethingmakenosenes")
	if user == nil {
		t.FailNow()
	}
	if !user.IsAnonymous() {
		t.Error("should be an anonymous user")
	}
}
