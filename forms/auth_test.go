package forms

import (
	r "github.com/christopherhesse/rethinkgo"
	"github.com/wangbin/imwb/settings"
	"testing"
)

func TestLoginForm(t *testing.T) {
	rs, _ := r.Connect(settings.DbUri, settings.DbName)
	form := NewLoginForm()
	if form.IsValid() {
		t.FailNow()
	}
	form.Name = "wangbin"
	form.Password = "1234"
	form.SetRs(rs)
	if !form.IsValid() {
		t.FailNow()
	}
}
