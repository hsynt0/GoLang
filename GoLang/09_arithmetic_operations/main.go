package main

import "fmt"

func main() {
	x, y := 15, 10

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", (x + y), (x + y))
	fmt.Printf("%T, %v\n", (x / y), (x / y))

	n := 5 / 2   // 2
	z := 5 / 2.0 // 2.5 or (5.0 / 2)
	fmt.Printf("z:%T, %v\n", z, z)
	fmt.Printf("n:%T, %v\n", n, n)

	fmt.Println("-------------------------")

	// Increment ++, Decrement --   Postfix
	// bunlar bir statement'dır. atamadır.
	h := 10

	h--

	fmt.Println(h)

	h++

	fmt.Println(h)

	h = h + 1

	fmt.Println(h) // fmt.Println(h++) yazdıramazsın

}
