package kata

import "math"

func SquareOrSquareRoot(arr []int) []int {
	var new_arr = []int{}

	for _, i := range arr {
		temp := math.Sqrt(float64(i))

		if temp == math.Trunc(temp) {
			new_arr = append(new_arr, int(temp))
		} else {
			temp = math.Pow(float64(i), 2)
			new_arr = append(new_arr, int(temp))
		}
	}

	return new_arr
}
