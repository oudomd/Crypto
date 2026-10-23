package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func stringToBinary(input string) string {
	binary := ""
	for _, char := range input {
		binary += fmt.Sprintf("%08b ", char)
	}
	return binary
}

func stringToHex(input string) string {
	return hex.EncodeToString([]byte(input))
}

func stringToBase64(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func main() {
	var input string
	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	binary := stringToBinary(input)
	hexadecimal := stringToHex(input)
	base64 := stringToBase64(input)

	fmt.Println("Binary representation:", binary)
	fmt.Println("Hexadecimal representation:", hexadecimal)
	fmt.Println("Base64 representation:", base64)
}
