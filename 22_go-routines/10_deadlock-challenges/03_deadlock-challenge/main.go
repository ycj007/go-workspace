package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	/*for  {
		fmt.Println(<-c)
	}*/

	select {
	case x := <-c:
		fmt.Println(x)

	}

}

// Why does this only print zero?
// And what can you do to get it to print all 0 - 9 numbers?
