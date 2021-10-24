package sha256

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

// Sha256String
func Sha256String(key, data []byte) string {
	return string(Sha256(key, data))
}

// Sha256Hex
func Sha256Hex(key, data []byte) string {
	return hex.EncodeToString(Sha256(key, data))
}

// Sha256
func Sha256(key, data []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(data)
	return mac.Sum(nil)
}
