package main

import "fmt"

func main() {

	const x = 5 // typeless

	var y float64 = 12.6

	fmt.Printf("%T, %v\n", x, x)

	fmt.Printf("%T, %v\n", y, y)

	fmt.Printf("%T, %v\n", x+y, x+y) // float64(x) + y

	fmt.Printf("%T, %v\n", x, x)

	const m = int16(5.2 + 4.8)

	fmt.Printf("%T, %v\t", m, m)

	const k = 3
	const l = 5.6 // eğer sen float ve int olarak belirtirsen hata verir.
	const r = true
	const s = "test"

	fmt.Printf("%T, %v\n", k, k)

	fmt.Printf("%T, %v\n", l, l)

	fmt.Printf("%T, %v\n", k+l, k+l)

	fmt.Printf("%T, %v\n", r, r)

	fmt.Printf("%T, %v\n", s, s)

}
