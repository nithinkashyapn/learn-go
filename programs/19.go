package main

import "fmt"

func main() {

	var arrayone = [5]int{1, 2, 3, 4, 5}
	var arraytwo = []int{0, 0, 0, 0, 0}

	for i := 0; i < len(arrayone); i++ {
		fmt.Println(arrayone[i])
	}

	for i := 0; i < len(arrayone); i++ {
		fmt.Println(arraytwo[i])
	}

}
