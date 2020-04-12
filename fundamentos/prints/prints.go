package main

import (
	"fmt"
)

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha.\n")

	x := 3.141516
	xs := fmt.Sprint(x)
	fmt.Println("O valor é " + xs)
	fmt.Println("O valor é", x)

	fmt.Printf("\nO Valor é %.2f", x)
}

