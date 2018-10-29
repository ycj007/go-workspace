package main

import "fmt"

func main() {

	varData := 10
	fmt.Println(varData)
	fmt.Printf("%d - %b \n", varData, varData)

	fmt.Printf("%d - %b - %x \n", varData, varData, varData)
	fmt.Printf("%d - %b - %#x \n", varData, varData, varData)
	fmt.Printf("%d - %b - %#X \n", varData, varData, varData)
	fmt.Printf("%d \t %b \t %#X \n", varData, varData, varData)

	for i := 0; i < 1000; i++ {
		fmt.Printf("%d \t %b \t %#x\n", i, i, i)

	}

}
