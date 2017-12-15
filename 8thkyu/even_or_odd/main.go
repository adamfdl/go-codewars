package main

import (
	"fmt"
)

func main() {
	result := EvenOrOdd(4)
	fmt.Println(result)
}

func EvenOrOdd(number int) bool {
	result := false
	if number%2 == 0 {
		result = true
	}
	return result
}
