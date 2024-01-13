package validator

import "unicode"

// check if password is complex enough
func isPasswordComplex(password string) bool {
	//check if the length is at least 8 characters
	if len(password) < 8 {
		return false
	}
	//check if the password contains at least one uppercase letter
	if !containsUppercase(password) {
		return false
	}
	//check if the password contains at least one lowercase letter
	if !containsLowercase(password) {
		return false
	}
	//check if the password contains at least one number
	if !containsNumber(password) {
		return false
	}
	return true

}
func containsUppercase(password string) bool {
	for _, char := range password {
		if unicode.IsUpper(char) {
			return true
		}
	}
	return false
}
func containsLowercase(password string) bool {
	for _, char := range password {
		if unicode.IsLower(char) {
			return true
		}
	}
	return false
}

func containsNumber(password string) bool {
	for _, char := range password {
		if unicode.IsNumber(char) {
			return true
		}
	}
	return false
}
