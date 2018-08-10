package main

import "fmt"

var a, b, c int = 20, 20, 20

func main() {

	fmt.Println("Before -- All globals")
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	a := 10
	b = 10

	fmt.Println("After -- ")
	fmt.Println("Local", a)
	fmt.Println("Global", b)
	fmt.Println("Global", c)

}
