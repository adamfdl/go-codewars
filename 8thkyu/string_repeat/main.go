package main

import (
	"bytes"
	"fmt"
)

func main() {
	result := RepeatStr(6, "I")
	fmt.Println(result)
}

func RepeatStr(repititions int, value string) string {
	var finalString bytes.Buffer
	for i := 0; i < repititions; i++ {
		finalString.WriteString(value)
	}
	return finalString.String()
}
