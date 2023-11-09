package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(Add(1, 2))
	fmt.Println(Add(2, 1))
}

func Add(a, b int) int {
	return a + b
}

func Sub(a, b int) int {
	return a - b
}
