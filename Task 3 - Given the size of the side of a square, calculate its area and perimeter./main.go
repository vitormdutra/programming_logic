package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Given the size of the side of a square, calculate its area and perimeter.")

	size := 3.0
	perimeter := size * 4
	area := math.Pow(size, 2)

	fmt.Printf("The perimeter is %.2f and area is %.2f", perimeter, area)
}
