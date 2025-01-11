package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Since the formula for converting Fahrenheit to Celsius is C = 5/9 (F – 32), read a temperature value in Fahrenheit and display it in Celsius")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type the temperature to converto do Celsius")
	temperature, _ := reader.ReadString('\n')
	temperature = strings.TrimSpace(temperature)
	temperatureConverted, _ := strconv.ParseFloat(temperature, 64)
	Celsius := (temperatureConverted - 32) / 1.8
	fmt.Printf("The temperature converted to celsius is %.2f", Celsius)
}
