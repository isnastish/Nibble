package validator

import (
	"regexp"
	"strings"
)

// Validate user's password. The password should contain at least 10,
// and at most 32 characters. It's a simplified version of how
// password validation could be achieved.
// Returns true if passowrd is valid, false otherwise.
func ValidateUserPassword(pwd string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9_@$#:&%]`)
	return len(pwd) >= 10 && len(pwd) <= 32 && re.MatchString(pwd)
}

// Validate user's email address. This is a bare minimum validation
// which should never be used in a production, but is sufficient for our purposes.
// Only checks if provided strings contains `@` symbol,
// if so, returns true, false otherwise.
// Since writing a complete email validator is not part of this assignment.
func ValidateUserEmailAddress(email string) bool {
	return len(email) >= 3 && len(email) <= 320 && strings.Contains(email, "@")
}
