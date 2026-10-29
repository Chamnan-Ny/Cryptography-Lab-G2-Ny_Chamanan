package main

import "fmt"


func xorEncrypt(text string, key byte) string {

	output := make([]byte, len(text))


	for i := 0; i < len(text); i++ {
		output[i] = text[i] ^ key
	}

	return string(output)
}

func main() {
	var message string
	fmt.Print("Enter a message to encrypt (no spaces): ")
	fmt.Scan(&message)

	var key byte = 'K'

	ciphertext := xorEncrypt(message, key)
	fmt.Println("Encrypted (Ciphertext):", ciphertext)
	fmt.Printf("Encrypted (in Hex):     %x\n", ciphertext)

	decrypted := xorEncrypt(ciphertext, key)
	fmt.Println("Decrypted (Original):  ", decrypted)
}