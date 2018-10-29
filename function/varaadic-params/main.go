package main

import "fmt"

func main() {

	n := avg(43, 56, 87, 12, 45, 57)
	datas := []int{1, 2, 3, 4, 5}
	fmt.Println(n)
	fmt.Print(avg(datas...))

}

func avg(datas ...int) int {

	var count int
	for _, v := range datas {

		count += v
	}
	return count / len(datas)

}
