package main

import (
	"fmt"
)

func main() {
	result := Opposite(-12)
	fmt.Println(result)
}

func Opposite(value int) int {
	return value * -1
}
