package sha

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"io"
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

func Sha1Hex(str string) (string, error) {
	b, err := Sha1(str)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

func Sha1(str string) ([]byte, error) {
	h := sha1.New()
	if _, err := io.WriteString(h, str); err != nil {
		return nil, err
	}

	return h.Sum(nil), nil
}
