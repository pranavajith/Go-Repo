package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Saying Hello to someone", func(t *testing.T) {
		got := Hello("Pranav", "")
		want := "Hello Pranav!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Say Hello World! when no name is given", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World!"
		assertCorrectMessage(t, got, want)
	})
	t.Run("Testing in Spanish to someone", func(t *testing.T) {
		got := Hello("Pranav", "hindi")
		want := "Namaste Pranav!"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q want %q", got, want)
	}
}
func TestAdd(t *testing.T) {
	t.Run("Adding numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15
		assertCorrectMessageInt(t, got, want, numbers)
	})
}

func assertCorrectMessageInt(t testing.TB, got int, want int, numbers []int) {
	t.Helper()
	if got != want {
		t.Errorf("Got %d want %d, %v", got, want, numbers)
	}
}
