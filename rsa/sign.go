package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
)

var (
	ErrBadPem = errors.New("密钥错误")
)

func Sign(data string, rsaPrivateKey []byte) (string, error) {
	block, _ := pem.Decode(rsaPrivateKey)
	if block == nil {
		return "", ErrBadPem
	}
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return "", err
	}

	hash := crypto.Hash.New(crypto.SHA1)
	hash.Write([]byte(data))

	sign, err := privKey.Sign(rand.Reader, hash.Sum(nil), crypto.SHA1)

	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(sign), nil
}

func CheckSign(data, target string, publicKey []byte) error {
	sign, err := base64.StdEncoding.DecodeString(target)
	if err != nil {
		return err
	}
	// public, _ := base64.StdEncoding.DecodeString(pubKey)

	block, _ := pem.Decode(publicKey)
	if block == nil {
		return ErrBadPem
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}
	pub := pubInterface.(*rsa.PublicKey)

	hash := sha1.New()
	hash.Write([]byte(data))
	return rsa.VerifyPKCS1v15(pub, crypto.SHA1, hash.Sum(nil), sign)
}
