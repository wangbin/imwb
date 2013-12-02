package hashers

import (
	"strings"
	"testing"
)

const (
	RawPass = "1qaz,2wsx"
)

var (
	ph256 = NewPBKDF2PasswordHash()
	ph1   = NewPBKDF2SHA1PasswordHash()
)

func TestSha256Encode(t *testing.T) {
	_, err := ph256.Encode(RawPass)
	if err != nil {
		t.Error(err)
	}
}

func TestSha256Verify(t *testing.T) {
	pw, _ := ph256.Encode(RawPass)
	if !ph256.Verify(RawPass, pw) {
		t.Error("Password should be verified")
	}
	if ph256.Verify("12345", pw) {
		t.Error("Password should not be verified")
	}
}

func TestSha256MustUpdate(t *testing.T) {
	pw, _ := ph256.Encode(RawPass)
	if ph256.MustUpdate(pw) {
		t.FailNow()
	}
	badPass := strings.Replace(pw, "12000", "12001", 1)
	if !ph256.MustUpdate(badPass) {
		t.FailNow()
	}
}

func TestSha1Encode(t *testing.T) {
	_, err := ph1.Encode(RawPass)
	if err != nil {
		t.Error(err)
	}
}

func TestSha1Verify(t *testing.T) {
	pw, _ := ph1.Encode(RawPass)
	if !ph1.Verify(RawPass, pw) {
		t.Error("Password should be verified")
	}
	if ph1.Verify("12345", pw) {
		t.Error("Password should not be verified")
	}
}

func TestSha1MustUpdate(t *testing.T) {
	pw, _ := ph1.Encode(RawPass)
	if ph1.MustUpdate(pw) {
		t.FailNow()
	}
	badPass := strings.Replace(pw, "12000", "12001", 1)
	if !ph1.MustUpdate(badPass) {
		t.FailNow()
	}
}
