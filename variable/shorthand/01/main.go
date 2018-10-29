package main

import (
	"fmt"
)

func main() {

	a := 1
	b := 1.23232
	c := "sfsfsf"
	d := 'a'
	e := `sfdsfsfsdfsfsfsf ?`
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)

}
