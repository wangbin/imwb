package hashers

import (
	"code.google.com/p/go.crypto/pbkdf2"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"
)

const (
	DefaultIterations = 12000
)

type PBKDF2PasswordHash struct {
	*BasePasswordHash
	iterations int
}

func NewPBKDF2PasswordHash() *PBKDF2PasswordHash {
	return &PBKDF2PasswordHash{&BasePasswordHash{
		algorithm: "pbkdf2_sha256", digest: sha256.New},
		DefaultIterations}
}

func NewPBKDF2SHA1PasswordHash() *PBKDF2PasswordHash {
	return &PBKDF2PasswordHash{&BasePasswordHash{
		algorithm: "pbkdf2_sha1", digest: sha1.New},
		DefaultIterations}
}

func (this *PBKDF2PasswordHash) salt() string {
	return RandomString()
}

func (this *PBKDF2PasswordHash) encodeWithIterations(password string,
	salt string, iterations int) (string, error) {
	keyLen := this.digest().Size()
	src := pbkdf2.Key([]byte(password), []byte(salt), iterations, keyLen,
		this.digest)
	dst := base64.StdEncoding.EncodeToString(src)
	h := strings.TrimSpace(string(dst))
	result := fmt.Sprintf("%s$%d$%s$%s", this.algorithm, iterations, salt, h)
	return result, nil
}

func (this *PBKDF2PasswordHash) Encode(password string) (string, error) {
	salt := this.salt()
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

func (this *PBKDF2PasswordHash) MustUpdate(encode string) bool {
	p := strings.SplitN(encode, "$", 4)
	if len(p) != 4 {
		return true
	}
	iterations, err := strconv.Atoi(p[1])
	if err != nil {
		return true
	}
	if iterations != this.iterations {
		return true
	}
	return false
}
