package main

import "fmt"

func main() {

	//greeting("ycj")
	//greetingWithTwoParam("yuan","chongjie")
	//fmt.Print(greetingWithReturn("chongjie","yuan"));

	fmt.Print(greetingWithTwoReturn("yuan", "chongjie"))
}

func greeting(name string) {
	fmt.Print(name)
}

func greetingWithTwoParam(firstName string, lastName string) {
	fmt.Print(firstName, lastName)
}

func greetingWithReturn(firstName string, lastName string) string {
	return fmt.Sprint(firstName, lastName)
}
func greetingWithoutReturn(firstName string, lastName string) string {
	/*	result:= fmt.Sprint(firstName,lastName)
	return */
	return ""
}

func greetingWithTwoReturn(firstName string, lastName string) (string, string) {
	return fmt.Sprint(firstName, lastName), fmt.Sprint(lastName, firstName)
}
