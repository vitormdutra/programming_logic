package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Given the size of the radius of a circle, calculate its area and perimeter.")

	radius := 5.0
	perimeter := 2 * math.Pi * radius
	area := math.Pi * (math.Pow(radius, 2))

	fmt.Printf("The area is %.2f and the perimeter is %.2f", area, perimeter)
}
