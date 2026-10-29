package main

import (
	"encoding/base64" 
	"encoding/hex"    
	"fmt"
)

func main() {
	var text string

	fmt.Print("Enter a word (no spaces): ")
	fmt.Scan(&text)

	fmt.Print("Binary:      ")
	for i := 0; i < len(text); i++ {
		fmt.Printf("%08b ", text[i])
	}
	fmt.Println() 

	hexResult := hex.EncodeToString([]byte(text))
	fmt.Println("Hexadecimal:", hexResult)

	base64Result := base64.StdEncoding.EncodeToString([]byte(text))
	fmt.Println("Base64:     ", base64Result)
}