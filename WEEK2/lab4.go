
package main

import "fmt"

func add(a int, b int) int { return a + b }
func sub(a int, b int) int { return a - b }
func mul(a int, b int) int { return a * b }
func div(a int, b int) int { return a / b }
func mod(a int, b int) int { return a % b }

func main() {

	for {
		fmt.Println("\n Mini Calculator ")
		fmt.Println("1) Add 2) Sub 3) Mul 4) Div 5) Mod 6) Exit")

		var choice int
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		if choice == 6 {
			fmt.Println("Exiting calculator. Goodbye!")
			break
		}

		if choice < 1 || choice > 6 {
			fmt.Println("Error: invalid operator choice. Please pick 1 to 6.")
			continue
		}

		var a, b int
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print("Enter b: ")
		fmt.Scan(&b)

		if choice == 1 {
			fmt.Println("Result:", add(a, b))
		} else if choice == 2 {
			fmt.Println("Result:", sub(a, b))
		} else if choice == 3 {
			fmt.Println("Result:", mul(a, b))
		} else if choice == 4 {
			if b == 0 {
				fmt.Println("Error: division by zero")
			} else {
				fmt.Println("Result:", div(a, b))
			}
		} else if choice == 5 {
			if b == 0 {
				fmt.Println("Error: division by zero")
			} else {
				fmt.Println("Result:", mod(a, b))
			}
		}
	}
}