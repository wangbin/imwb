package hashers

import (
	"code.google.com/p/go.crypto/bcrypt"
	"crypto/sha256"
	"fmt"
)

type BcryptPasswordHash struct {
	*BasePasswordHash
}

func NewBcryptPasswordHash() *BcryptPasswordHash {
	return &BcryptPasswordHash{&BasePasswordHash{algorithm: "bcrypt"}}
}

func NewBcryptSHA256PasswordHash() *BcryptPasswordHash {
	return &BcryptPasswordHash{&BasePasswordHash{algorithm: "bcrypt_sha256",
		digest: sha256.New}}
}

func (this *BcryptPasswordHash) Encode(password string) (string, error) {
	pw := []byte(password)
	if this.digest != nil {
		h := this.digest()
		h.Write(pw)
		pw = h.Sum(nil)
	}

	encode, err := bcrypt.GenerateFromPassword(pw, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s$%s", this.algorithm, encode), nil
}

func (this *BcryptPasswordHash) Verify(password, encode string) bool {
	encode = string(encode[len(this.algorithm)+1:])
	pw := []byte(password)
	if this.digest != nil {
		h := this.digest()
		h.Write(pw)
		pw = h.Sum(nil)
	}
	err := bcrypt.CompareHashAndPassword([]byte(encode), pw)
	if err != nil {
		return false
	}
	return true
}
