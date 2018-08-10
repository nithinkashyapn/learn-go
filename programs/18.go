package main

import "fmt"
import "strings"

func main() {

	exampleString := "ABCDEF"
	fmt.Println("Length is ", len(exampleString))

	exampleArray := []string{"A", "B", "C", "D"}
	fmt.Println(strings.Join(exampleArray, " "))
}
