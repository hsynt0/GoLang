package main

import "fmt"

func main() {

	// var name string = "Hüseyin"
	/*var(
		name = "Hüseyin"
		age int16 = -256
		isMarried bool = true
		weight float32 = 72.5
	)
	*/
	// var name, age, isMarried, weight = "Hüseyin", 25, true, 72.5
	name, age, isMarried, weight := "Hüseyin", 25, true, 72.5
	var height int
	// bool ---> false, numeric ---> 0, string --->""

	fmt.Println(height)
	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isMarried)
	fmt.Println(weight)
}
