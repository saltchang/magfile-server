package util

import (
	"crypto/sha512"
	"encoding/base64"
)

func GetPasswordHash(password, salt string) string {
	passwordBytes := []byte(password)
	saltBytes := []byte(salt)
	passwordBytes = append(passwordBytes, saltBytes...)

	sha512 := sha512.New()
	sha512.Write(passwordBytes)
	checkSum := sha512.Sum(nil)

	checkSumBase64 := base64.URLEncoding.EncodeToString(checkSum)

	return checkSumBase64
}

func CheckPasswordHash(password, salt, hash string) bool {
	checkSum := GetPasswordHash(password, salt)
	return checkSum == hash
}
