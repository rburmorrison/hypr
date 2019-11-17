package html

import "testing"

func assertExpectation(t *testing.T, expected, actual string) {
	t.Errorf("unexpected output, expected %q, got %q", expected, actual)
}
