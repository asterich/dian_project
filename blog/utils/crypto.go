package utils

import (
	"blog/utils/errmsg"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func ValidatePassword(hashedPassword string, plainPassword string) (bool, errmsg.ErrCode) {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword)); err != nil {
		return false, errmsg.ERROR_PASSWORD_INCORRECT
	}
	return true, errmsg.SUCCEED
}

func GeneratePassword(plainPassword string) (string, errmsg.ErrCode) {
	var hashedPassword, err = bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Failed to generate hashed password, err: ", err)
		return string(hashedPassword), errmsg.ERROR
	}
	return string(hashedPassword), errmsg.SUCCEED
}
