package main

import "testing"

func TestAdd(t *testing.T) {
	if total := Add(5, 5); total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
