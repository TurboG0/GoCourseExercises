package main

import (
	"testing"
)

func TestVtestaob(t *testing.T) {
	if got, want := vtestaob("Ushashvelo"),
		"Funqciis argumentia: Ushashvelo"; got != want {
			t.Errorf("error - want %s got %s", want, got)
	}
}
