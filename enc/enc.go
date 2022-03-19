package enc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/hex"

	"golang.org/x/crypto/pbkdf2"
)

func DeriveKey(passphrase string, salt []byte) []byte {
	return pbkdf2.Key([]byte(passphrase), salt, 4096, 32, sha1.New)
}

func Encrypt(plaintext string, iv []byte) (string, error) {
	key := GetKey()

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	plaintextBytes := []byte(plaintext)
	cfb := cipher.NewCFBEncrypter(block, iv)
	ciphertext := make([]byte, len(plaintextBytes))
	cfb.XORKeyStream(ciphertext, plaintextBytes)

	encoded := hex.EncodeToString(ciphertext)

	return encoded, nil
}

func Decrypt(ciphertext string, iv []byte) (string, error) {
	key := GetKey()

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	decoded, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	cfb := cipher.NewCFBDecrypter(block, iv)
	plaintext := make([]byte, len(decoded))
	cfb.XORKeyStream(plaintext, decoded)

	return string(plaintext), nil
}
