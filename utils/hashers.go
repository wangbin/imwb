package utils

import (
	"code.google.com/p/go.crypto/pbkdf2"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"hash"
	"strconv"
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
	digest     func() hash.Hash
}

func NewPBKDF2PasswordHash() *PBKDF2PasswordHash {
	return &PBKDF2PasswordHash{iterations: 12000, algorithm: "pbkdf2_sha256", digest: sha256.New}
}

func (this *PBKDF2PasswordHash) Salt() string {
	return RandomString()
}

func (this *PBKDF2PasswordHash) encodeWithIterations(password string, salt string, iterations int) (string, error) {
	keyLen := this.digest().Size()
	src := pbkdf2.Key([]byte(password), []byte(salt), iterations, keyLen,
		this.digest)
	dst := base64.StdEncoding.EncodeToString(src)
	h := strings.TrimSpace(string(dst))
	result := fmt.Sprintf("%s$%d$%s$%s", this.algorithm, iterations, salt, h)
	return result, nil
}

func (this *PBKDF2PasswordHash) Encode(password string, salt string) (string, error) {
	return this.encodeWithIterations(password, salt, this.iterations)
}

func (this *PBKDF2PasswordHash) Verify(password, encode string) bool {
	p := strings.SplitN(encode, "$", 4)
	if len(p) != 4 {
		return false
	}
	algorithm, salt := p[0], p[2]
	iterations, err := strconv.Atoi(p[1])
	if err != nil {
		return false
	}
	if algorithm != this.algorithm {
		return false
	}
	encode2, err := this.encodeWithIterations(password, salt, iterations)
	if err != nil {
		return false
	}
	if len(encode) != len(encode2) {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(encode), []byte(encode2)) == 1
}
