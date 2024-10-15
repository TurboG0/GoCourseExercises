package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if got, want := add(6, 5), 11; got != want {
		t.Errorf("error - want %d got %d", want, got)
	}
}

func TestSub(t *testing.T) {
	if got, want := sub(19, 12), 7; got != want {
		t.Errorf("error - want %d got %d", want, got)
	}
}

func TestDoMath(t *testing.T) {
	if got, want := add(6, 5), 11; got != want {
		t.Errorf("error - want %d got %d", want, got)
	}

	if got, want := sub(19, 12), 7; got != want {
		t.Errorf("error - wamt %d got %d", want, got)
	}
}
