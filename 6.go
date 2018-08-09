package main

import "fmt"

func main() {
	const length float32 = 10.0
	const width float32 = 20.0

	area := length * width

	fmt.Printf("%v * %v is %f\n", length, width, area)
}
