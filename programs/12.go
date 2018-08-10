package main

import "fmt"

func compare(int1 int, int2 int) int {
	if int1 < int2 {
		return int2
	} else {
		return int1
	}
}

func main() {

	a := compare(10, 20)
	b := compare(30, 20)

	fmt.Println(a)
	fmt.Println(b)
}
