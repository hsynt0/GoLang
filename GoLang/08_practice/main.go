package main

import (
	"fmt"
	"strconv"
)

func main() {

	x := 4
	var y float64

	y = float64(x) // type(value)
	fmt.Println(y)
	/*
	 */
	a := 10
	b := 5
	println("a:", a, "b:", b)
	a, b = b, a // a = b, b = a
	println("a:", a, "b:", b)
	/*
	 */
	yaş := 40
	fmt.Println(yaş) // You can use turkish character or someone else
	心臓 := "Gönül"
	fmt.Println(心臓)
	/*
	 SHADOWİNG */
	z := 3
	if true {
		z := 6
		fmt.Println(z)
	}
	fmt.Println(z)
	//////////////////////////
	m := 10
	if true {
		m = 6
		m++
		fmt.Println(m)
	}
	fmt.Println(m)
	/*
	 */
	n := (65)

	s := string(n)

	fmt.Printf("%v, %T\n", n, n) // 65
	fmt.Printf("%v, %T\n", s, s) // A

	p := strconv.Itoa(n)
	fmt.Printf("%v, %T\n", p, p) // "65" string olarak 65 yazdırdı.

}
