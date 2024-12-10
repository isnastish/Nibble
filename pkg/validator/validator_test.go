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
}
