package main

import "fmt"

func main() {

	/*	for i:=0;i<10;i++ {

			fmt.Print(i)

		}

		for i :=0;i<10;i++ {
			for j:=i;j<10;j++{
				fmt.Printf("%d --%d \n",i,j)
			}
		}*/

	h := 10
	for h <
		100 {
		fmt.Printf("%d \n", h)
		h++
	}

	for {
		h++
		fmt.Printf("%d \n", h)
		if h > 200 {
			break
		}

	}

	for {
		h++
		if h%2 == 0 {
			continue
		}
		fmt.Println(h)
		if h >= 50 {
			break
		}
	}

	for i := 250; i <= 340; i++ {
		fmt.Println(i, " - ", string(i), " - ", []byte(string(i)))
	}
	foo := "a"
	fmt.Println(foo)
	fmt.Printf("%T \n", foo)

}
