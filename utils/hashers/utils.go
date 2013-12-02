package hashers

import (
	"crypto/rand"
	"math/big"
)

const (
	AllowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	MaxSaltSize  = 12
)

func RandomString() string {
	result := make([]byte, MaxSaltSize)
	length := int64(len(AllowedChars))
	for index := range result {
		i, _ := rand.Int(rand.Reader, big.NewInt(length))
		result[index] = AllowedChars[i.Int64()]
	}
	return string(result)
}
