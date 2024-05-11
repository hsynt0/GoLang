package main

import "fmt"

//  packVar := "Package Scope" olmaz.
var packVar = "Package Scope" // Paket dışındaki değişkenler hafızada gereksiz yer kaplayABİLİR.
var funcVar = "Func(Package) Scope"

func main() {

	if true {
		var blockVar = "Block Scope"
		fmt.Println(blockVar)
	}
	var funcVar = "Func Scope"

	fmt.Println(funcVar)

	fmt.Println(packVar)

	myFunc()
	////////////////////
	////////////////////////////
	///////////////////
	var name = "Hüseyin"
	fmt.Println(name)
	name, surName := "Kamil", "Miras"
	fmt.Println(name, surName)

}

func myFunc() {
	fmt.Println(funcVar)
}
