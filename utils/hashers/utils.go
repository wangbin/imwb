package hashers

import (
	"math/rand"
)

const (
	AllowedChars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func RandomString() string {
	result := make([]byte, 12)
	length := len(AllowedChars)
	for index := range result {
		result[index] = AllowedChars[rand.Intn(length)]
	}
	return string(result)
}
