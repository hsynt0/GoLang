package main

import "fmt"

func main() {

	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	// Type Conversion tyoe(value) => int(y) = 10.0 =>10

	println(x + int(y)) // y  değişkeninin int sadece kısmını aldı.
	fmt.Printf("%v %T\n", y, y)

	/////////////////////////////////////////////////////////////////

	/* var a int8 = 10
	   var b int16 = 10

	fmt.Println(a+b) yapamazsın. int olsa da değişkeni farklı */

	//String ifadesi de int bir değere dönüştürülemez. Şimdilik

	num1 := 106
	str1 := string(num1)

	fmt.Printf("%v %T\n", num1, num1)
	fmt.Println()
	fmt.Printf("%v %T\n", str1, str1)

}
