package main

import (
	"fmt"
)

func average(array []int, size int) float64 {

	sum := 0
	for i := 0; i < size; i++ {
		sum += array[i]
	}
	return float64(sum) / float64(size)

}

func main() {

	var array = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var avg = average(array, len(array))
	fmt.Printf("%f\n", avg)

}
