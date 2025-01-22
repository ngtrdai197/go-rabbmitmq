package helper

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func GenerateRandomStr() string {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
}

func RandomString(prefix string) string {
	if prefix != "" {
		return fmt.Sprintf("%s%s", prefix, GenerateRandomStr())
	}
	return GenerateRandomStr()
}
