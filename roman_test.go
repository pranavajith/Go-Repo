package main

import "testing"

func TestRomanNumeras(t *testing.T) {
	got := ConvertToRoman(1984)
	want := "MCMLXXXIV"
	if got != want {
		t.Errorf("Got %s want %s", got, want)
	}
}
