package utils

import (
	"testing"
)

func TestSha256(t *testing.T) {
	hash := Sha256([]byte("my-password@"))
	// Hash is take from third-party service generated with sha256 algorithm
	if hash != "f3827ee3d1896ccb409e80d1b30de197f848b3223ccfa0d466dd77d5e7f24f7a" {
		t.Errorf("hash mismatch")
	}

	// Hash of an empty byte steam
	hash = Sha256([]byte{})
	if hash != "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" {
		t.Errorf("hash mismatch")
	}
}
