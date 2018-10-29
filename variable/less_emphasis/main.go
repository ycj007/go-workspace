package main

import "fmt"

var vTest8 = "vTest8"

func main() {

	/**
	  declear var
	*/

	var vTest = 10
	var vTest2 int = 11
	vTest3 := 12

	var vTest4, vTest5 = 13, false
	var vTest6, vTest7 string
	vTest6 = "15-str"
	vTest8 := false
	fmt.Println(vTest, vTest2, vTest3, vTest4, vTest5, vTest6, vTest7, vTest8)
	name := `Todd` // back-ticks work like double-quotes
	fmt.Println("Hello ", name)

}
