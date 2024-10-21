package lib

import (
	"crypto/sha1"
	"encoding/hex"
)

func Encrypt(keyString string) string {
	hasher := sha1.New()
	hasher.Write([]byte(keyString))
	return hex.EncodeToString(hasher.Sum(nil))
}
