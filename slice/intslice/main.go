package main

import "fmt"

func main() {
	ops()
}

func ops() {

	mySlice := make([]int, 0, 3)

	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice))
	fmt.Println(mySlice)
	fmt.Printf("%T \n", mySlice)

	for i := 0; i < 10; i++ {

		mySlice := append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value: ", mySlice)
	}

}

func intSlice() {
	mySlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("%T \n", mySlice)
	fmt.Println(mySlice)
}

func forLoop() {
	xs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, value := range xs {
		fmt.Println(i, " - ", value)
	}
}
