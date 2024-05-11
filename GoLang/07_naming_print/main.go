package main

import "fmt"

var B = 100

func main() {

	// (Print - Println) - Printf

	fmt.Println("hi guys") //Println yazıldıktan sonra satır atlanır. Kendi satırını bitirir.
	fmt.Print("Hi Guys")
	fmt.Println("sa")
	fmt.Print("Merhaba Arkadaşlar! ")

	name := "Hüso"

	fmt.Print(name)
	fmt.Println(name)
	fmt.Printf(name)

	fmt.Println()

	fmt.Print("Benim Adim:", name)
	fmt.Println("")
	fmt.Println("Benim Adim:", name)                     // String ifadeden sonra boşluk koydu
	fmt.Printf("Benim Adim: %v %T %X", name, name, name) // "" %v = variable yani değişken %T = Type yani türü golang.org/pkg/fmt/

	fmt.Println("\n///////////////////////////")

	x := 100
	y := 20
	z := 30

	fmt.Printf("%b %d %o", x, y, z) // golang.org/pkg/fmt/
	fmt.Println("\n/////////////////////////////")

	surName, age := "Turco", 19

	fmt.Print("Benim Adim ", name, surName, " ve ben ", age, " yasindayim")
	fmt.Println("\nBenim Adim", name, surName, "ve ben", age, "yasindayim") // Look at the space
	fmt.Printf("Benim Adim %v %v ve ben %v yasindayim", name, surName, age)

	fmt.Println("\n/////////////////////////////")

	// VISIBILITY

	/* var coin string
	// Go da camelcase meşhur morqu
	var coinType string
	// Kısaltmalar büyük harfle genelde i, j, k
	var URL
	var HTTP
	*/
}
