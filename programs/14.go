package main

import "fmt"
import "math"

func main() {

	SquareRoot := func(x float64) float64 {
		return math.Sqrt(x)
	}

	fmt.Println(SquareRoot(16))
}
