package utils

import "golang.org/x/crypto/bcrypt"

func GetPasswordHash(password []byte) (string, error) {
	hash, err := bcrypt.GenerateFromPassword(password, 0)
	return string(hash), err
}
