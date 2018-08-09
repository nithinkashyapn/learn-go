package main

import "fmt"

type Square struct {
	length, width int
}

func area(tray Square) int {
	return tray.length * tray.width
}

func main() {

	tray := Square{length: 10, width: 10}
	fmt.Printf("Area is %v\n", area(tray))

}
