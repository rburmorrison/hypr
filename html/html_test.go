package html

import "testing"

func assertExpectation(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("unexpected output, expected %q, got %q", expected, actual)
	}
}
