package utils

import "golang.org/x/crypto/bcrypt"

func CheckPasswordHash(password, hash string) bool {
	if len(password) == 0 || len(hash) == 0 {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	}

	return true
}
