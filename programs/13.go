package main

import "fmt"

func reverse(x, y, z string) (string, string, string) {
	return z, y, x
}
func main() {
	a, b, c := reverse("A", "B", "C")
	fmt.Println(a, b, c)
}
