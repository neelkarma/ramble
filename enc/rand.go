package enc

import "crypto/rand"

func GenerateIv() []byte {
	iv := make([]byte, 16)
	rand.Read(iv)
	return iv
}

func GenerateSalt() []byte {
	salt := make([]byte, 8)
	rand.Read(salt)
	return salt
}
