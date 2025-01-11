package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Read an integer and display its successor.")
	read := bufio.NewReader(os.Stdin)
	fmt.Println("Type a integer")
	integer, _ := read.ReadString('\n')
	integer = strings.TrimSpace(integer)
	i, _ := strconv.Atoi(integer)
	fmt.Printf("Integer typed is: %d", i)
	fmt.Println()
	fmt.Printf("Your sucessor is: %d", i+1)
}
