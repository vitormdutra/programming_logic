package main

import (
	"bufio"
	"fmt"
	"math"
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

func convertKhToMs(max_speed int) float64 {
	conversion := float64(max_speed) / 3.6
	return conversion
}

func meterToGetMaxSpeed(max_speed float64) float64 {
	s := 1 * 2.5 * math.Pow(max_speed, 2) / 2
	return s
}

func main() {
	fmt.Println("In a city, it is desired to synchronize traffic lights. In this way, when a traffic light turns green, vehicles that were stopped at it tend to find the next traffic lights also open. In order for this to be done, the next traffic lights need to open a little later, depending on the speed limit on the road and the distance between them. Thus, when the traffic light turns green, a vehicle starts to accelerate until it reaches the permitted speed, which it maintains until it reaches the next traffic light, taking a certain amount of time to travel this distance. In order for it to find the next traffic light open, it must open a little before the vehicle arrives (for example: 3 seconds before). Therefore, create an algorithm that informs how long after a traffic light should open, given the following information: a. the distance from the previous traffic light b. the permitted speed of the road c. the typical acceleration of cars")
	traffic_light_green_time := 30.0
	fmt.Println("Using a 1.0 car we can move 2.5m/s, type the max speed of the road")
	meterPerSecond := keyboardInput()
	meterPerSecond = convertKhToMs(int(meterPerSecond))
	timeToMaxSpeed := meterPerSecond / 2.5
	fmt.Println()
	fmt.Printf("need this time to get max speed limit %.2f", timeToMaxSpeed)
	fmt.Println()
	fmt.Printf("meter in start to max speed %.2f", meterToGetMaxSpeed(timeToMaxSpeed))
	fmt.Println()

	fmt.Println("Type the distance to next light traffic")
	distanceToLight := keyboardInput()
	timeToFinish := distanceToLight - meterToGetMaxSpeed(timeToMaxSpeed)
	resultAtMaxSpeed := timeToFinish / meterPerSecond
	resultTime := resultAtMaxSpeed + timeToMaxSpeed
	fmt.Println()
	fmt.Printf("This is the time to go in 500 meter %.2f", resultTime)
	fmt.Println()

	if resultTime > traffic_light_green_time {
		result := traffic_light_green_time - resultTime
		if result < 0 {
			result = math.Abs(result)
			fmt.Printf("The traffic light need to have this time to car get in gree %.2f", result)
		} else {
			fmt.Printf("The traffic light need to have this time to car get in gree %.2f", result)
		}
	}
}
