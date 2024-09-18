package main

import (
	"fmt"
	"math"
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
	// t.Run("Adding numbers", func(t *testing.T) {
	// 	numbers := []int{1, 2, 3, 4, 5}
	// 	got := Sum(numbers)
	// 	want := 15
	// 	assertCorrectMessageInt(t, got, want, numbers)
	// })
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

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		repeat("a")
	}
}

func TestValues(t *testing.T) {

	t.Run("Testing Simple R Perimeter", func(t *testing.T) {
		got := perimeter(10.0, 10.0)
		want := 40.0
		if got != want {
			t.Errorf("Got %.2f want %.2f", got, want)
		}
	})
	t.Run("Testing R Area", func(t *testing.T) {
		got := area(10.0, 20.0)
		want := 200.0
		if got != want {
			t.Errorf("Got %.2f want %.2f", got, want)
		}
	})
	r := Rectangle{10.0, 10.0}
	t.Run("Perimeter of Rectangle struct", func(t *testing.T) {
		got := perimRect(r)
		want := 40.0
		assertCorrectMessageInt(t, got, want)

	})
	t.Run("Area of Rectangle struct", func(t *testing.T) {
		got := areaRect(r)
		want := 100.0
		assertCorrectMessageInt(t, got, want)

	})
	c := Circle{10.0}
	t.Run("Perimeter of Circle struct", func(t *testing.T) {
		got := perimCircle(c)
		want := math.Pi * 10.0 * 2
		assertCorrectMessageInt(t, got, want)

	})
	t.Run("Area of Rectangle struct", func(t *testing.T) {
		got := areaCircle(c)
		want := math.Pi * 100.0
		assertCorrectMessageInt(t, got, want)

	})
}

func assertCorrectMessageInt(t testing.TB, got float64, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("Got %g want %g", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	c := Circle{10.0}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		if want != shape.Area() {
			t.Errorf("Got %.2f want %.2f", shape.Area(), want)
		}
	}
	t.Run("Checking area method of Rectangle", func(t *testing.T) {
		checkArea(t, r, 100.0)
	})
	t.Run("Checking area method of Circle", func(t *testing.T) {
		checkArea(t, c, 100.0*math.Pi)
	})
}

func TestWallet(t *testing.T) {

	t.Run("Testing Balance function", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("Error")
		}
	})
	t.Run("Testing Deposit function", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Deposit(Bitcoin(20))
		got := wallet.Balance()
		want := Bitcoin(30)
		if got != want {
			t.Errorf("Error. Got %s, want %s", got, want)
		}
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()

		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("withdraw error", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		er := wallet.Withdraw(Bitcoin(30))

		// got := wallet.Balance()

		// want := Bitcoin(10)

		fmt.Printf("%s", er)

		if er == nil {
			t.Error("got no error, wanted one")
		}
	})

}
