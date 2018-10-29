package main

import "fmt"

func main() {
	test3()

}

func test3() {
	//var x [256]int;
	var x [256]string
	//var x [256]byte;
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 0; i < 256; i++ {

		x[i] = string(i)
	}

	for index, value := range x {

		fmt.Printf("%v-%T-%b \n", value, value, value)
		if index > 50 {
			break
		}

	}
}
func test2() {

	var x [100]string
	for i := 65; i < 123; i++ {
		x[i-65] = string(i)

	}
	fmt.Println(x)
	fmt.Print(x[42])

}
func test01() {
	var x [50]int
	fmt.Println(x)
	fmt.Println(x[40])
	x[40] = 100
	fmt.Println(x[40])
}
