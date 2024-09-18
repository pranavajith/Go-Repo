package main

import (
	"math"
)

// 17-09

// func main() {
// 	fmt.Println("Hello, world")
// }

//

func Hello(name string, language string) string {

	curPrefix := "Hello "
	curSuffix := "!"
	hindi := "hindi"
	if name == "" {
		name = "World"
	}
	if language == hindi {
		curPrefix = "Namaste "
	}
	return curPrefix + name + curSuffix
}

func Sum(number []int) int {
	ans := 0
	// for i := 0; i < len(number); i++ {
	// 	ans += number[i]
	// }
	for _, num := range number {
		ans += num
	}
	return ans
}

func SumAll(allSlices ...[]int) []int {
	sum := make([]int, len(allSlices))
	for i, slice := range allSlices {
		sum[i] = Sum(slice)
	}
	return sum
}

func SumTails(allSlices ...[]int) []int {
	var sum []int
	for _, slice := range allSlices {
		if len(slice) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(slice[1:]))
		}
	}
	return sum
}

func repeat(char string) string {
	num := 5
	var ans string
	for i := 0; i < num; i++ {
		ans += char
	}
	return ans
}

func perimeter(width float64, height float64) float64 {
	return 2 * (width + height)
}

func area(width float64, height float64) float64 {
	return width * height
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func perimRect(rect Rectangle) float64 {
	return (rect.Width + rect.Height) * 2
}
func areaRect(rect Rectangle) float64 {
	return (rect.Width * rect.Height)
}
func perimCircle(c Circle) float64 {
	return 2 * math.Pi * c.Radius
}
func areaCircle(c Circle) float64 {
	return math.Pi * c.Radius * c.Radius
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (c Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}
