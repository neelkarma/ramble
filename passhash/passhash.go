package passhash

import "golang.org/x/crypto/bcrypt"

func HashPassphrase(passphrase string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(passphrase), 14)
	return string(bytes), err
}

func ValidatePassphrase(passphrase string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passphrase))
	return err == nil
}
