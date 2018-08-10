package main

import (
	"fmt"
)

func main() {
	exampleSlice := make([]int, 3, 10)

	exampleSlice[0] = 1
	exampleSlice[1] = 2
	exampleSlice[2] = 3

	fmt.Printf("Slice is %v capacity is %v and length is %v\n", exampleSlice, cap(exampleSlice), len(exampleSlice))

	exampleSlice = append(exampleSlice, 4)
	exampleSlice = append(exampleSlice, 5)

	fmt.Printf("Slice is %v capacity is %v and length is %v\n", exampleSlice, cap(exampleSlice), len(exampleSlice))

	temp := make([]int, len(exampleSlice), len(exampleSlice))
	copy(temp, exampleSlice)
	fmt.Printf("Temp is %v capacity is %v and length is %v\n", temp, cap(temp), len(temp))

	fmt.Println("exampleSlice[2:4] is ", exampleSlice[2:4])
	fmt.Println("exampleSlice[2:] is ", exampleSlice[2:])
	fmt.Println("exampleSlice[:4] is ", exampleSlice[:4])

	twoDSlice := make([][]int, 3)
	for i := 0; i < 3; i++ {
		twoDSlice[i] = make([]int, i+1)
		for j := 0; j < i+1; j++ {
			twoDSlice[i][j] = i + j
		}
	}

	fmt.Println("Two Dimensional Slice is ", twoDSlice)

	fmt.Println("On expanding")

	for i := 0; i < 3; i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("[%v][%v] is %v\t", i, j, twoDSlice[i][j])
		}
		fmt.Println()
	}
}
