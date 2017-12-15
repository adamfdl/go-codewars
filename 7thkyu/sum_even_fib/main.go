package main

import (
	"fmt"
)

func main() {
	result := SumEvenFibonacci(7347000000)
	fmt.Println(result)
}

func SumEvenFibonacci(limit int) int {
	fibSeq := make([]int, limit+1)

	fibSeq[0] = 0
	fibSeq[1] = 1

	var sumEven int
	for i := 2; i <= limit; i++ {

		fibSeq[i] = fibSeq[i-1] + fibSeq[i-2]
		if fibSeq[i]%2 == 0 {
			sumEven += fibSeq[i]
		}

		if fibSeq[i] > limit {
			break
		}

	}

	return sumEven
}
