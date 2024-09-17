package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Prana")
	want := "Hello Pranav!"

	if got != want {
		t.Errorf("got %q wanted %q", got, want)
	}
}
