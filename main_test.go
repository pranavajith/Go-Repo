package main

import (
	"reflect"
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
	t.Run("testing adders", func(t *testing.T) {
		got := SumAll([]int{1, 2}, []int{3, 4, 5, 6}, []int{3, 4, 5, 6})
		want := []int{3, 18, 18}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Wanted %v got %v", want, got)
		}
	})
	t.Run("Adding all but first of all of them", func(t *testing.T) {
		got := SumTails([]int{}, []int{1, 2, 3}, []int{2, 3, 4, 5})
		want := []int{0, 5, 12}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v want %v", got, want)
		}
	})
}

func assertCorrectMessageInt(t testing.TB, got int, want int, numbers []int) {
	t.Helper()
	if len(numbers) == 0 {
		numbers = []int{0}
	}
	if got != want {
		t.Errorf("Got %d want %d, %v", got, want, numbers)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a")
	}
}
