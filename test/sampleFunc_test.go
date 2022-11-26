package sampleTest

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}

func TestMultiply(t *testing.T) {
	total := Multiply(5, 5)
	if total != 25 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
