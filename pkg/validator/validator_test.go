package validator

import (
	"testing"
)

func TestPasswordValidation(t *testing.T) {
	pwd := ""
	if ValidateUserPassword(pwd) { // empty password
		t.Errorf("failed to validate password %s", pwd)
	}

	pwd = "hello"
	if ValidateUserPassword(pwd) { // too short
		t.Errorf("failed to validate password %s", pwd)
	}

	pwd = "!!!!!not allowed"
	if ValidateUserPassword(pwd) { // contains not allowed symbols
		t.Errorf("failed to validate password %s", pwd)
	}

	pwd = "isnastish@1234"
	if !ValidateUserPassword(pwd) {
		t.Errorf("password validation should have passed %s", pwd)
	}

	pwd = "too long of a password, so that the validation should fail"
	if ValidateUserPassword(pwd) {
		t.Errorf("failed to validation the password %s", pwd)
	}
}

func TestEmailValidation(t *testing.T) {
	email := ""
	if ValidateUserEmailAddress(email) {
		t.Errorf("failed to validate email address %s", email)
	}

	email = "@"
	if ValidateUserEmailAddress(email) { // doesn't contain domain name
		t.Errorf("failed to validate email address %s", email)
	}

	email = "isnastish@gmail.com"
	if !ValidateUserEmailAddress(email) {
		t.Errorf("email validation should have passed %s", email)
	}
}
