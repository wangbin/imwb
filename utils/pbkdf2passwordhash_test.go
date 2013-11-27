package utils

import (
	"testing"
)

const (
	RawPass = "1qaz,2wsx"
	Salt    = "wangbin"
	Pass    = "pbkdf2_sha256$12000$wangbin$vMDHyaC5vsRWHgU2WvUuq1RCTYNVxte+3SR0+WINQ7Y="
)

func TestEncode(t *testing.T) {
	ph := NewPBKDF2PasswordHash()
	pw, err := ph.Encode(RawPass, Salt)
	if err != nil {
		t.Error(err)
	}
	if pw != Pass {
		t.Error(pw)
	}
}

func TestVerify(t *testing.T) {
	ph := NewPBKDF2PasswordHash()
	if !ph.Verify(RawPass, Pass) {
		t.Error("Password should be verified")
	}
	if ph.Verify("12345", Pass) {
		t.Error("Password should not be verified")
	}
}
