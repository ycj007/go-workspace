package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {

	data := 12
	/**
	showing
	*/

	fmt.Print(&data)
	/**
	using

	*/

	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")

}
