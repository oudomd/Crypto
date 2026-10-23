package main

import "fmt"

func main() {
	var num1, num2 int 

	fmt.Print("Enter first interger: ")
	fmt.Scan(&num1)
	fmt.Print("Enter second interger: ")
	fmt.Scan(&num2)

	bothPositve := num1 > 0 && num2 > 0
	oneGreater := num1 > num2 || num2 > num1 

	fmt.Println("Both numbers are positive:", bothPositve)
	fmt.Println("One number is greater than the other:", oneGreater)
}
