package revenuecat

import "testing"

func TestError(t *testing.T) {
	err := Error{
		Code:    123,
		Message: "Error message",
	}
	if str := err.Error(); str != "123: Error message" {
		t.Errorf("got: %q, expected: %q", str, "123: Error message")
	}
}
