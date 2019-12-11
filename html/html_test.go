package html

import "testing"

func assertExpectation(t *testing.T, expected, actual string) {
	if expected != actual {
		t.Errorf("unexpected output. got %q, expected %q", actual, expected)
	}
}
