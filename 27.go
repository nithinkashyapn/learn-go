package main

import (
	"fmt"
)

func main() {
	squareTables := make(map[int]int)

	squareTables[1] = 1
	squareTables[2] = 4
	squareTables[3] = 9
	squareTables[4] = 16

	fmt.Println(squareTables)

	delete(squareTables, 4)
	fmt.Println(squareTables)

	delete(squareTables, 3)
	fmt.Println(squareTables)

}
