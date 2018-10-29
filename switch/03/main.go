package main

import "fmt"

func main() {

	switch 1 {
	case 1, 2, 3:
		fmt.Print("1-3")
	case 4, 5, 6:
		fmt.Print("4-6")

	default:
		fmt.Print("0")

	}

}
