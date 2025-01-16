package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Convert an integer less than 32 to its binary representation")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type a number between 0 and 32")
	number, _ := reader.ReadString('\n')
	number = strings.TrimSpace(number)
	convertedNumber, _ := strconv.Atoi(number)
	if convertedNumber > 32 {
		log.Fatal("Number > 32")
	} else if convertedNumber < 0 {
		log.Fatal("Number < 0")
	} else {
		fmt.Println(strconv.FormatInt(int64(convertedNumber), 2))
	}
}
