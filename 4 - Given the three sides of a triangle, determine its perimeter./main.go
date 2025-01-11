package main

import "fmt"

var t1, t2, t3 float64

func main() {
	fmt.Println("Given the three sides of a triangle, determine its perimeter.")
	t1 = 3
	t2 = 3
	t3 = 3
	result := t1 + t2 + t3

	fmt.Printf("The perimeter is %.2f", result)
}
