package utils

import (
	"crypto/sha256"
	"fmt"
)

// Produce a hash using sha256 algorithm from a sequence of bytes
func Sha256(bytes []byte) string {
	hash := sha256.New()
	hash.Write(bytes)
	b := hash.Sum(nil)
	return fmt.Sprintf("%x", b)
}
