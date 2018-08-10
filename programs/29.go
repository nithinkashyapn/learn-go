package main

import "fmt"

type Compute interface {
	area() int
}

type Square struct {
	length int
}

type Rectangle struct {
	length, width int
}

func (square Square) area() int {
	return square.length * square.length
}

func (rectangle Rectangle) area() int {
	return rectangle.length * rectangle.width
}

func getArea(object Compute) int {
	return object.area()
}

func main() {

	square := Square{length: 10}
	rectangle := Rectangle{length: 5, width: 7}

	fmt.Printf("Area of square is %v\n", getArea(square))
	fmt.Printf("Area of rectangle is %v\n", getArea(rectangle))

}
