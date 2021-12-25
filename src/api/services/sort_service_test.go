package services

import "testing"

func TestConstants(t *testing.T) {
	// We can only test privateConst from the same package
	if privateConst != "private" {
		t.Error("privateConst should be 'private'")
	}
}
