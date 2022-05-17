package validation

import (
	"regexp"
)

const (
	UsernameOk uint = iota
	UsernameTooShort
	UsernameTooLong
	UsernameInvalidCharacter
)

const (
	PasswordOk uint = iota
	PasswordTooShort
	PasswordTooLong
	PasswordInvalidCharacter
)

func ValidateUsername(username string) uint {
	if len(username) < 1 {
		return PasswordTooShort
	}
	if len(username) > 24 {
		return PasswordTooLong
	}
	if match, _ := regexp.MatchString("^[a-zA-Z0-9-_.]+$", username); !match {
		return PasswordInvalidCharacter
	}
	return PasswordOk
}

func ValidatePassword(password string) uint {
	if len(password) < 8 {
		return PasswordTooShort
	}
	if len(password) > 48 {
		return PasswordTooLong
	}
	if match, _ := regexp.MatchString("^[a-zA-Z0-9-_@#$,.]+$", password); !match {
		return PasswordInvalidCharacter
	}
	return PasswordOk
}
