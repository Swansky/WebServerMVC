package utils

import (
	"awesomeProject1/models"
	"crypto/md5"
	"encoding/hex"
)

func ComparePassword(password string, user *models.User) bool {
	passwordHash := HashPassword(password)
	passwordMatch := passwordHash == user.Password
	if passwordMatch {
		return true
	}
	return false
}

func HashPassword(password string) string {
	passwordHash := md5.Sum([]byte(password))
	return hex.EncodeToString(passwordHash[:])
}
