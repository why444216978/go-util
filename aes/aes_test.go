package aes

import (
	"crypto/aes"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
	"gopkg.in/go-playground/assert.v1"
)

var key = []byte("1234567890abcdef")

func TestECB(t *testing.T) {
	Convey("TestECB", t, func() {
		Convey("success", func() {
			data := "123"
			res, err := EncryptECB([]byte(data), key)
			assert.Equal(t, err, nil)
			res, err = DecryptECB(res, key)
			assert.Equal(t, res, []byte(data))
		})
		Convey("EncryptECB key length error", func() {
			key := "123"
			res, err := EncryptECB([]byte("123"), []byte(key))
			assert.Equal(t, err, aes.KeySizeError(len(key)))
			assert.Equal(t, res, nil)
		})
		Convey("DecryptECB key length error", func() {
			key := "123"
			res, err := DecryptECB([]byte("123"), []byte(key))
			assert.Equal(t, err, aes.KeySizeError(len(key)))
			assert.Equal(t, res, nil)
		})
	})
}

func TestCBC(t *testing.T) {
	Convey("TestCBC", t, func() {
		Convey("success", func() {
			res, err := EncryptCBC([]byte("123"), key)
			assert.Equal(t, err, nil)
			res, err = DecryptCBC(res, key)
			assert.Equal(t, res, []byte("123"))
		})
		Convey("EncryptCBC key length error", func() {
			key := "123"
			res, err := EncryptCBC([]byte("123"), []byte("123"))
			assert.Equal(t, err, aes.KeySizeError(len(key)))
			assert.Equal(t, res, nil)
		})
		Convey("DecryptCBC key length error", func() {
			key := "123"
			res, err := DecryptCBC([]byte("123"), []byte(key))
			assert.Equal(t, err, aes.KeySizeError(len(key)))
			assert.Equal(t, res, nil)
		})
	})
}
