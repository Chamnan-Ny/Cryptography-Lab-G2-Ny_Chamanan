package main

import "fmt"

func myXor(a int, b int) {
	fmt.Println("Bitwise XOR (a ^ b):", a^b)
}

func myNOT(a int, b int) {
	fmt.Println("Bitwise NOT (^a):", ^a, "| (^b):", ^b)
}

func myOR(a int, b int) {
	fmt.Println("Bitwise OR (a | b):", a|b)
}

func myAND(a int, b int) {
	fmt.Println("Bitwise AND (a & b):", a&b)
}

func myShifts(a int, b int) {
	fmt.Println("Left Shift (a << 1):", a<<1)
	fmt.Println("Right Shift (a >> 1):", a>>1)
}

func main() {
	var a int
	var b int

	fmt.Print("Enter first integer (a): ")
	fmt.Scan(&a)
	fmt.Print("Enter second integer (b): ")
	fmt.Scan(&b)

	fmt.Println("\nBitwise Function Results")
	myXor(a, b)
	myNOT(a, b)
	myOR(a, b)
	myAND(a, b)
	myShifts(a, b)

	fmt.Println("\nEffect of Bitwise Assignment Operators")
	a &= b 
	fmt.Println("After a &= b, new value of a is:", a)

	a |= b 
	fmt.Println("After a |= b, new value of a is:", a)

	a ^= b 
	fmt.Println("After a ^= b, new value of a is:", a)
}
