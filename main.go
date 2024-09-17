package main

import "fmt"

// 17-09

// func main() {
// 	fmt.Println("Hello, world")
// }

//

func Hello(name string) string {
	if name != "" {
		return "Hello " + name + "!"
	}
	return "Hello World!"
}

func main() {
	fmt.Println()
}
