package main

import "fmt"

func main() {
	fmt.Println("An electrical circuit is composed of two resistors R1 and R2 in parallel, and both in sequence of a resistor R3. Create an algorithm to calculate the equivalent resistance of this circuit.")

	r1 := 23
	r2 := 46
	r3 := 96

	r12 := (r1 * r2) / (r1 + r2)
	req := r12 + r3
	fmt.Printf("Result = %d", req)
}
