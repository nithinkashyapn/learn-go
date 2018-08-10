package main

import "fmt"

func main() {

	var number = 100

	a := number << 2
	b := number >> 2
	c := number & 001
	d := number ^ 001
	e := number | 001

	fmt.Printf("%v after doing << is %v\n", number, a)
	fmt.Printf("%v after doing >> is %v\n", number, b)
	fmt.Printf("%v after doing & is %v\n", number, c)
	fmt.Printf("%v after doing ^ is %v\n", number, d)
	fmt.Printf("%v after doing | is %v\n", number, e)

}
