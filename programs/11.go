package main

import "fmt"

func main() {

	var examplearray = [10]int{0}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	for i := 0; i < 10; {
		fmt.Println(i)
		i += 3
	}

	for i := range examplearray {
		fmt.Println(i)
	}

}
