package main

import "fmt"

func Add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(Add(2, 3))
	fmt.Println(Add(-2, -8))
}
