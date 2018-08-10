package main

import (
	"fmt"
)

func main() {
	var a int = 100
	var b *int
	var c **int

	b = &a
	c = &b

	fmt.Printf("A is %v &A is %v\n", a, &a)
	fmt.Printf("B is %v *B is %v &B is %v\n", b, *b, &b)
	fmt.Printf("C is %v *C is %v &C is %v\n", c, *c, &c)
}
