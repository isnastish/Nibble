package validator

import (
	"regexp"
)

// Validate user's password. The password should contain at least 12,
// and at most 32 characters. It's a simplified version of how
// password validation could be achieved.
// Returns true if passowrd is valid, false otherwise.
func ValidateUserPassword(pwd string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_@$#:&%]{12,32}$`)
	return re.MatchString(pwd)
}

func ValidateUserEmailAddress(email string) bool {
	return false
}
