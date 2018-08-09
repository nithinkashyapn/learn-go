package main

import "fmt"

func main() {
	var a, b, c = 10, 20.0, "hello"

	fmt.Printf("%v is of type %T\n", a, a)
	fmt.Printf("%v is of type %T\n", b, b)
	fmt.Printf("%v is of type %T\n", c, c)
}
