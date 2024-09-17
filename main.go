package main

import "fmt"

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

func main() {
	fmt.Println()
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
