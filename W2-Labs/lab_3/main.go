package main 

import "fmt"

func myXor(a, b int) {
	result := a ^ b 
	fmt.Printf("XOR of %d and %d is : %d\n", a, b, result)
}

func myNOT(a, b int) {
	resultA := ^a
	resultB := ^b 
	fmt.Printf("NOT of %d is: %d\n", a, resultA)
	fmt.Printf("NOT of %d is: %d\n", b, resultB)
}

func myOR(a, b int) {
	result := a | b 
	fmt.Printf("OR of %d and %d is: %d\n", a, b, result)
}

func myAND(a, b int) {
	result := a << 1 
	fmt.Printf("AND of %d and %d is: %d\n", a, b, result)
}

func myLeftShift(a, b int) {
	resultA := a << 1
	resultB := b << 1
	fmt.Printf("Left shift of %d is: %d\n", a, resultA)
	fmt.Printf("Left shift of %d is: %d\n", b, resultB)
}

func myRightShift(a, b int) {
	resultA := a >> 1
	resultB := b >> 1
	fmt.Printf("Right shift of %d is: %d\n", a, resultA)
	fmt.Printf("Right shift of %d is: %d\n", b, resultB)
}

func assignmentOperatorsDemo(a, b int) {
	fmt.Printf("Initial values: a = %d, b = %d\n", a, b)
	a += b
	fmt.Printf("After a += b, a = %d\n", a)
	a -= b
	fmt.Printf("After a -= b, a = %d\n", a)
	a *= b
	fmt.Printf("After a *= b, a = %d\n", a)
	if b != 0 {
		a /= b
		fmt.Printf("After a /= b, a = %d\n", a)
	}
	a %= b
	fmt.Printf("After a %%= b, a = %d\n", a)
}

func main() {
	var a, b int
	fmt.Print("Enter first integer: ")
	fmt.Scan(&a)
	fmt.Print("Enter second integer: ")
	fmt.Scan(&b)

	myXor(a, b)
	myNOT(a, b)
	myOR(a, b)
	myAND(a, b)
	myLeftShift(a, b)
	myRightShift(a, b)
	assignmentOperatorsDemo(a, b)
}

