package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Request a person's age in days and display the age in years, months and days on the screen.")

	fmt.Println("Type your age")
	reader := bufio.NewReader(os.Stdin)
	age, _ := reader.ReadString('\n')
	age = strings.TrimSpace(age)
	ageConverted, _ := strconv.Atoi(age)
	years := ageConverted
	months := ageConverted * 12
	days := ageConverted * 365
	fmt.Printf("Your age is %d, this in months is %d, and this in days is %d", years, months, days)
}
