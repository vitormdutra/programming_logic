package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertMsToKh() float64 {
	reader := bufio.NewReader(os.Stdin)
	meters, _ := reader.ReadString('\n')
	meters = strings.TrimSpace(meters)
	convertMeters, _ := strconv.ParseFloat(meters, 64)
	kilometers := convertMeters * 3.6
	return kilometers
}

func main() {
	fmt.Println("Create an algorithm that transforms a speed provided in m/s by the user into km/h. To do this, multiply the value in m/s by 3.6.")
	fmt.Println("Type the meters you want to convert to kilometers")
	result := convertMsToKh()
	fmt.Printf("The conversion is %.2f", result)
}
