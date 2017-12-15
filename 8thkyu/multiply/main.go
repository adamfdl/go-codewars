package main

import (
	"fmt"
)

func main() {
	result := Multiply(5, 4)
	fmt.Println(result)
}

func Multiply(a, b int) int {
	return a * b
}
