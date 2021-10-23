package crypto_utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateToken(length uint) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
