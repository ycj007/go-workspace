package main

import (
	"fmt"
	"strings"
)

func makeGreeter() func() string {
	test := func() int {
		return 12
	}
	return
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
	fmt.Printf("%T\n", greet)
}
