package main

import (
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	return a / b, nil
}

func modulus(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("modulus by zero")
	}
	return a % b, nil
}

func main() {
	for {
		fmt.Println("===== Mini Calculator =====")
		fmt.Println("1) Add")
		fmt.Println("2) Sub")
		fmt.Println("3) Mul")
		fmt.Println("4) Div")
		fmt.Println("5) Mod")
		fmt.Println("6) Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Exiting...")
			break
		}

		var a, b int
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		switch choice {
		case 1:
			fmt.Printf("Result: %d\n", add(a, b))
		case 2:
			fmt.Printf("Result: %d\n", subtract(a, b))
		case 3:
			fmt.Printf("Result: %d\n", multiply(a, b))
		case 4:
			result, err := divide(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %d\n", result)
			}
		case 5:
			result, err := modulus(a, b)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("Result: %d\n", result)
			}
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
