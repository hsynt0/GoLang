package main

import (
	"fmt"
	"math"
)

func main() {

	/* 5
	3.14
	"passed"  constant
	true
	*/
	r := 3.0

	const pi float64 = 3.14

	areaOfCircle := pi * (math.Pow(r, 2))

	fmt.Println(areaOfCircle)

	const x int = 2
	const y float32 = 3.4
	const z string = "test"
	const t bool = false

	fmt.Printf("%T, %v\n", x, x)
	fmt.Printf("%T, %v\n", y, y)
	fmt.Printf("%T, %v\n", z, z)
	fmt.Printf("%T, %v\n", t, t)

	/* const n = 2

	n = 5, x++ you cannot bcs thats const */

	// const ---> compile time
	// var ---> run time

	/* var m = math.Pow(3, 4)
	const m = math.Pow(3, 4) değerini compile edlilirken bilinmesi gerek fakat math fonksiyoonu run timeda oluşur
	fmt.Printf("%T, %v\n", m, m)
	*/

	k := 3
	// const p = k  run timeda oluşacak bir değeri constant bir değere atayamıyoruz.
	const p = 1
	const u, o = 3, 5

	fmt.Printf("%T, %v\n", k, k)
	fmt.Printf("%T, %v\n", p, p)
	fmt.Printf("%T, %v\n", k+p, k+p)

}
