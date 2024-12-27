package utils

import (
	"regexp"
)

var (
	uppercaseRegex   = regexp.MustCompile(`[A-Z]`)
	lowercaseRegex   = regexp.MustCompile(`[a-z]`)
	digitRegex       = regexp.MustCompile(`\d`)
	specialCharRegex = regexp.MustCompile(`[#?!@$%^&*-]`)
	minLength        = 8
	maxLength        = 32
)

/*
IsValidPassword reports whether the password is valid.

A valid password must contain:
  - Minimum eight characters,
  - At most thirty-two characters,
  - At least one uppercase letter,
  - At least one lowercase letter,
  - At least one number,
  - At least one special character.
*/
func IsValidPassword(password string) bool {
	if len(password) < minLength || len(password) > maxLength {
		return false
	}
	if !uppercaseRegex.MatchString(password) {
		return false
	}
	if !lowercaseRegex.MatchString(password) {
		return false
	}
	if !digitRegex.MatchString(password) {
		return false
	}
	if !specialCharRegex.MatchString(password) {
		return false
	}
	return true
}
