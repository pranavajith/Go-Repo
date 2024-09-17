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
