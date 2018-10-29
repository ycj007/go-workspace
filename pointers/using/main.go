package main

import "fmt"

func zero(z int) {
	fmt.Printf("%p\n", &z) // address in func zero
	fmt.Println(&z)        // address in func zero
	z = 0
}

func zero2(z *int) {
	fmt.Println(z)
	*z = 0
}
func main() {
	x := 5
	fmt.Printf("%p\n", &x) // address in main
	fmt.Println(&x)        // address in main
	zero(x)
	zero2(&x)
	fmt.Println(x) // x is still 5
}
