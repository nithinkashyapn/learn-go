package main

import (
	"fmt"
	"time"
)

func recursion() {
	time.Sleep(1000)
	fmt.Println("Recursion")
	main()
}

func main() {
	time.Sleep(1000)
	fmt.Println("Main")
	recursion()
}
