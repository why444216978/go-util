package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
)

// PublicEncrypt 公钥加密
func PublicEncrypt(encryptStr string, publicKey []byte) (result string, err error) {
	// pem 解码
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return "", ErrBadPem
	}

	// x509 解码
	publicKeyInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return
	}
	pub := publicKeyInterface.(*rsa.PublicKey)

	// 对明文进行加密
	encryptedStr, err := rsa.EncryptPKCS1v15(rand.Reader, pub, []byte(encryptStr))
	if err != nil {
		return
	}

	// 返回密文
	return base64.URLEncoding.EncodeToString(encryptedStr), nil
}

// PrivateDecrypt 私钥解密
func PrivateDecrypt(decryptStr string, privateKey []byte) (result string, err error) {
	decryptBytes, err := base64.URLEncoding.DecodeString(decryptStr)
	if err != nil {
		return
	}

	// pem 解码
	block, _ := pem.Decode(privateKey)

	// X509-PKCS8 解码
	_priv, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return
	}
	priv := _priv.(*rsa.PrivateKey)

	// 对密文进行解密
	decrypted, err := rsa.DecryptPKCS1v15(rand.Reader, priv, decryptBytes)
	if err != nil {
		return
	}

	// 返回明文
	return string(decrypted), nil
}
