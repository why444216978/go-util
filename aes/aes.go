package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"errors"
)

func EncryptECB(plaintext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plaintext = pkcsPadding(plaintext, blockSize)

	ciphertext := make([]byte, len(plaintext))
	for bs, be := 0, blockSize; bs < len(plaintext); bs, be = bs+blockSize, be+blockSize {
		block.Encrypt(ciphertext[bs:be], plaintext[bs:be])
	}

	return ciphertext, nil
}

func DecryptECB(ciphertext, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	if len(ciphertext)%blockSize != 0 {
		return nil, errors.New("密文长度不是块大小的整数倍")
	}

	plaintext := make([]byte, len(ciphertext))
	for bs, be := 0, blockSize; bs < len(ciphertext); bs, be = bs+blockSize, be+blockSize {
		block.Decrypt(plaintext[bs:be], ciphertext[bs:be])
	}

	return pkcsUnPadding(plaintext), nil
}

func EncryptCBC(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	data = pkcsPadding(data, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	encrypted := make([]byte, len(data))
	blockMode.CryptBlocks(encrypted, data)
	return encrypted, nil
}

func DecryptCBC(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	decrypted := make([]byte, len(data))
	blockMode.CryptBlocks(decrypted, data)
	decrypted = pkcsUnPadding(decrypted)
	return decrypted, nil
}

func pkcsPadding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

func pkcsUnPadding(decrypted []byte) []byte {
	length := len(decrypted)
	unPadding := int(decrypted[length-1])
	return decrypted[:(length - unPadding)]
}
