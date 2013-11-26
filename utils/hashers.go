package utils

import (
	"code.google.com/p/go.crypto/pbkdf2"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

type PasswordHasher interface {
	Salt() string
	Verify(password, encode string) bool
	Encode(password, salt string) string
	SafeSummery(encode string) map[string]string
	MustUpdate() bool
}

type PBKDF2PasswordHash struct {
	iterations int
	algorithm  string
}

func NewPBKDF2PasswordHash() *PBKDF2PasswordHash {
	return &PBKDF2PasswordHash{12000, "pbkdf2_sha256"}
}

func (this *PBKDF2PasswordHash) Salt() string {
	return RandomString()
}

func (this *PBKDF2PasswordHash) Encode(password string, salt string) (string, error) {
	src := pbkdf2.Key([]byte(password), []byte(salt), this.iterations, 32,
		sha256.New)
	dst := base64.StdEncoding.EncodeToString(src)
	hash := strings.TrimSpace(string(dst))
	result := fmt.Sprintf("%s$%d$%s$%s", this.algorithm, this.iterations, salt,
		hash)
	return result, nil
}
