/*
  Package tools
  @Author: Ahsen17
  @Github: https://github.com/Ahsen17
  @Time: 2024/6/2 13:34
  @Description: ...
*/

package tools

import (
	"crypto/aes"
	"encoding/hex"
)

var salt = "blogserv@github.com//AhsenEdward"

// encryptAES AES加密
func encryptAES(key string, plainText string) (string, error) {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	out := make([]byte, len(plainText))
	cipher.Encrypt(out, []byte(plainText))

	return hex.EncodeToString(out), nil
}

// decryptAES AES解密
func decryptAES(key string, encryptText string) (string, error) {
	cipher, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	decodeText, _ := hex.DecodeString(encryptText)
	out := make([]byte, len(decodeText))
	cipher.Decrypt(out, decodeText)

	return string(out[:]), nil
}

// EncryptAES 加密
func EncryptAES(plainText string) (string, error) {
	return encryptAES(salt, plainText)
}

// DecryptAES 解密
func DecryptAES(encryptText string) (string, error) {
	return decryptAES(salt, encryptText)
}
