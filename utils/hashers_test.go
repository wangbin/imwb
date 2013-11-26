package utils

import (
	"testing"
)

func TestPBKDF2PasswordHashEncode(t *testing.T) {
	ph := NewPBKDF2PasswordHash()
	pw, err := ph.Encode("1qaz,2wsx", "wangbin")
	if err != nil {
		t.Error(err)
	}
	if pw != "pbkdf2_sha256$12000$wangbin$vMDHyaC5vsRWHgU2WvUuq1RCTYNVxte+3SR0+WINQ7Y=" {
		t.Error(pw)
	}
}
