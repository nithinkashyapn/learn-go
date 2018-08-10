package main

import "fmt"

func main() {

	var a = 10
	var b = 20
	var c, d = 30, 40

	if b < a {
		fmt.Printf("A < B\n")
	} else {

		fmt.Printf("In else part\n")

		if b == 20 {
			fmt.Printf("Nested If\n")
		}

		switch c {
		case 10:
			fmt.Println("Switch 10")
		case 20:
			fmt.Println("Switch 20")
		case 30:
			fmt.Println("Switch 30")
		default:
			fmt.Println("Default")
		}

	}

	fmt.Printf("D is %v\n", d)
}
