package hashers

import (
	"hash"
)

type PasswordHasher interface {
	Verify(password, encode string) bool
	Encode(password string) string
	MustUpdate(encode string) bool
}

type BasePasswordHash struct {
	algorithm string
	digest    func() hash.Hash
}
