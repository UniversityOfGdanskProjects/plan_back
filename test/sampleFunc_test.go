package sampleTest

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 12 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
