package main

import (
	"fmt"
)

func main() {

	a := []int{1, 2, 3, 4, 5}
	var pointer [5]*int

	for i := 0; i < 5; i++ {
		pointer[i] = &a[i]
	}

	for i := 0; i < 5; i++ {
		fmt.Printf("%d -- Pointer = %d -- *Pointer = %d\n", i, pointer[i], *pointer[i])
	}

}
