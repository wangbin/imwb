package hashers

import (
	"hash"
)

var (
	Algorithm string
	hasherMap = make(map[string]PasswordHasher)
)

func init() {
	hasherMap["pbkdf2_sha256"] = NewPBKDF2PasswordHash()
	hasherMap["pbkdf2_sha1"] = NewPBKDF2SHA1PasswordHash()
	hasherMap["bcrypt"] = NewBcryptPasswordHash()
	hasherMap["bcrypt_sha256"] = NewBcryptSHA256PasswordHash()
	Algorithm = "bcrypt_sha256"
}

type PasswordHasher interface {
	Verify(password, encode string) bool
	Encode(password string) (string, error)
}

type BasePasswordHash struct {
	algorithm string
	digest    func() hash.Hash
}

func MakePassword(rawPass string) string {
	p, err := hasherMap[Algorithm].Encode(rawPass)
	if err != nil {
		return ""
	}
	return p
}

func CheckPassword(rawPass, encode string) bool {
	return hasherMap[Algorithm].Verify(rawPass, encode)
}
