package main

import (
	"fmt"
)

func xorEncryptDecrypt(text string, key byte) string {
	result := make([]byte, len(text))
	for i := 0; i < len(text); i++ {
		result[i] = text[i] ^ key
	}
	return string(result)
}

func main() {
	var plaintext string
	var key byte

	fmt.Print("Enter plaintext: ")
	fmt.Scanln(&plaintext)
	fmt.Print("Enter key (single byte as an integer): ")
	fmt.Scanf("%d", &key)

	ciphertext := xorEncryptDecrypt(plaintext, key)
	fmt.Println("Ciphertext:", ciphertext)

	decryptedText := xorEncryptDecrypt(ciphertext, key)
	fmt.Println("Decrypted text:", decryptedText)
}
