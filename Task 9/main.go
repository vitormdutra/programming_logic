package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("9 - Create an algorithm that calculates and presents the value of the volume of an oil can, given its radius and height.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type the radius of you oil can")
	radius, _ := reader.ReadString('\n')
	radius = strings.TrimSpace(radius)
	convertedRadius, _ := strconv.ParseFloat(radius, 64)
	fmt.Println("Type the height of your oil can")
	height, _ := reader.ReadString('\n')
	height = strings.TrimSpace(height)
	convertedHeight, _ := strconv.ParseFloat(height, 64)
	volume := math.Pi * convertedRadius * convertedHeight
	fmt.Printf("The volume of your oil can is %.2f", volume)
}
