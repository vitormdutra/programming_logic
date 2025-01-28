package main

import "fmt"

func main() {
	fmt.Println("A program to manage withdrawals from an ATM must have some mechanism to decide the number of notes of each value that should be made available to the customer who made the withdrawal. One possible criterion would be the 'optimal distribution' criterion, in the sense that the lowest value notes should be distributed in the minimum possible number. For example, if the amount requested was R$87.00, the program should indicate one R$50.00 note, three R$10.00 notes, one R$5.00 note and two R$1.00 notes. Write a program that receives the value of the amount requested and returns the distribution of the notes according to the optimal distribution criterion (consider there are notes of R$1.00; R$2.00; R$5.00; R$10.00; R$20.00; R$50.00 and R$100.00).")

	cashSolicitated := 78
	distributed := cashSolicitated
	for i := 0; distributed > 0; i++ {
		if (distributed / 100) >= 1 {
			fmt.Println("send 100$")
			distributed = distributed - 100
		} else if (distributed / 50) >= 1 {
			fmt.Println("send 50$")
			distributed = distributed - 50
		} else if (distributed / 20) >= 1 {
			fmt.Println("send 20$")
			distributed = distributed - 20
		} else if (distributed / 10) >= 1 {
			fmt.Println("send 10$")
			distributed = distributed - 10
		} else if (distributed / 5) >= 1 {
			fmt.Println("send 5$")
			distributed = distributed - 5
		} else if (distributed / 2) >= 1 {
			fmt.Println("send 2$")
			distributed = distributed - 2
		} else if (distributed / 1) >= 1 {
			fmt.Println("send 1$")
			distributed = distributed - 1
		}
	}
}
