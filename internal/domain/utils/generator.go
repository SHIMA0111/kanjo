package utils

import (
	"crypto/rand"
	"encoding/hex"
)

// RandomString generates a random hexadecimal string of the specified length.
// The function uses cryptographic randomness for secure generation.
func RandomString(length int) string {
	resultBytes := make([]byte, length)

	if _, err := rand.Read(resultBytes); err != nil {
		// From the crypto/rand manual contract us that never return error
		panic(err)
	}

	randomString := hex.EncodeToString(resultBytes)

	return randomString
}
