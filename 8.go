package main

import "fmt"

func main() {

	var a = 100
	var b = &a
	var c = *b
	var d = &b
	var e = &c
	var f = *d
	var g = *e

	fmt.Printf("a i.e, value is %v\n", a)
	fmt.Printf("b i.e., address of a is %v\n", b)
	fmt.Printf("c i.e., value of b is %v\n", c)
	fmt.Printf("d i.e., address of b is %v\n", d)
	fmt.Printf("e i.e., address of c is %v\n", e)
	fmt.Printf("f i.e., value of d (address of b) is %v\n", f)
	fmt.Printf("g i.e., value of e (address of c) is %v\n", g)
}
