package main

import "fmt"

func increment() func() int {
	var i int = 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	variable1 := increment()
	variable2 := increment()

	fmt.Println(variable1())
	fmt.Println(variable1())
	fmt.Println(variable2())
	fmt.Println(variable2())
}
