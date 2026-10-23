package main

import "fmt"

func main() {
	var a, b int

	fmt.Print("Enter first interger: ")
	fmt.Scan(&a)
	fmt.Print("Enter Second interger: ")
	fmt.Scan(&b)
	
	fmt.Printf("Initial values: a = %d, b = %d\n", a, b)

	a = b 
	fmt.Printf("After a = b, a = %d\n", a)
	
	a += b 
	fmt.Printf("After a += b , a = %d\n", a)

	a -= b 
	fmt.Printf("After a -= b , a = %d\n", a)

	a *= b
	fmt.Printf("After a *= b , a = %d\n", a)

	if b != 0 {
		a /= b
		fmt.Printf("After a /= b , a = %d\n", a)

		a %= b
		fmt.Printf("After a %%= b , a = %d\n", a)
	} else { 
		fmt.Println("Division and modulus operations skipped as b is 0.")
	}
}
