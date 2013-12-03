package hashers

import (
	"testing"
)

var (
	b256 = NewBcryptSHA256PasswordHash()
	b1   = NewBcryptPasswordHash()
)

func TestBcryptSHA256Encode(t *testing.T) {
	_, err := b256.Encode(RawPass)
	if err != nil {
		t.Error(err)
	}
}

func TestBcryptSHA256Verify(t *testing.T) {
	pw, _ := b256.Encode(RawPass)
	if !b256.Verify(RawPass, pw) {
		t.Error("Password should be verified")
	}
	if b256.Verify("12345", pw) {
		t.Error("Password should not be verified")
	}
}

func TestBcryptEncode(t *testing.T) {
	_, err := b1.Encode(RawPass)
	if err != nil {
		t.Error(err)
	}
}

func TestBcryptVerify(t *testing.T) {
	pw, _ := b1.Encode(RawPass)
	if !b1.Verify(RawPass, pw) {
		t.Error("Password should be verified")
	}
	if b1.Verify("12345", pw) {
		t.Error("Password should not be verified")
	}
}
