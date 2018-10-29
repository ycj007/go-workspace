package main

import "fmt"

func main() {

	data := 100

	fmt.Print(data)

	dataPointer := &data

	*dataPointer = 223213
	fmt.Print(*dataPointer)
}
