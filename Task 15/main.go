package main

import "fmt"

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func main() {
	fmt.Println("Write a program to generate the inverse of a three-digit number (example: the inverse of 498 is 894).")
	fmt.Println(reverse("543"))
}
