package main

import "fmt"

func main() {

	if true {
		fmt.Print("true")
	}

	if false {
		fmt.Println("false" +
			"")
	}

	if !true {

		fmt.Println("!true")
	}

	if !false {
		fmt.Println("!false")
	}
	if true || false {

		fmt.Print("true || false")
	}

	if true && false {

		fmt.Print("true&&false")
	}

}
