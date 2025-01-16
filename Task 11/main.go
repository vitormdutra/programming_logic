package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func keyboardInput() float64 {
	reader := bufio.NewReader(os.Stdin)
	word, _ := reader.ReadString('\n')
	word = strings.TrimSpace(word)
	convertedResult, _ := strconv.ParseFloat(word, 64)
	return convertedResult
}

func main() {
	fmt.Println("Create an algorithm to calculate a student's semester grade. The semester grade is obtained by the arithmetic mean of the grades of 2 semesters. Each semester grade is composed of 2 test grades.")
	fmt.Println("Type the n1 for b1")
	b1n1 := keyboardInput()
	fmt.Println("Type the n2 for b1")
	b1n2 := keyboardInput()
	resultB1 := (b1n1 + b1n2)
	resultB1 = resultB1 / 2

	fmt.Println("Type the n1 for b2")
	b2n1 := keyboardInput()
	fmt.Println("Type the n2 for b2")
	b2n2 := keyboardInput()
	resultB2 := (b2n1 + b2n2)
	resultB2 = resultB2 / 2

	resultSemester := (resultB1 + resultB2) / 2
	fmt.Printf("%.2f", resultSemester)
}
