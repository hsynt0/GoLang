package main

import "fmt"

const sabit = 14

func main() {

	x := 50
	x = x - 10 // Assignment Statement
	x -= 10    // Assignment Operation

	fmt.Printf("%T, %v\n", x, x)

	// K = F - 32 / 1.8 + 273 => -49 F kaç K derecedir?

	F := -40
	K := float64(F-32)/1.8 + 273

	fmt.Printf("%T, %v\n", K, K)

	//	Shadowing sabitlerde de geçerli midir? -Evet.

	const sabit = 32
	fmt.Printf("Sabit: %T, %v\n", sabit, sabit)

}
