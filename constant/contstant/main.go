package main

import "fmt"

const p = "print p"

const (
	a = "a"
	b = "b"
)

func main() {

	const q = "print q"

	//p= "sfsf"
	fmt.Print(p, q)
	fmt.Print(a, b)
}
