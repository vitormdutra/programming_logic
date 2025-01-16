package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Read two integers and display the quotient and remainder of the integer division between them")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Type the first number")
	n1, _ := reader.ReadString('\n')
	n1 = strings.TrimSpace(n1)
	fmt.Println("type the second number")
	n2, _ := reader.ReadString('\n')
	n2 = strings.TrimSpace(n2)
	cn1, _ := strconv.Atoi(n1)
	cn2, _ := strconv.Atoi(n2)

	quotient := cn1 / cn2
	remainder := cn1 % cn2

	fmt.Printf("Quotient: %d\n", quotient)
	fmt.Printf("Remainder: %d\n", remainder)
}
