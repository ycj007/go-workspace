package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}
	for key, val := range myGreeting {
		fmt.Println(key, " - ", val)
	}
	/*	fmt.Println(myGreeting)

		if val, exists := myGreeting[7]; exists {
			delete(myGreeting, 7)
			fmt.Println("val: ", val)
			fmt.Println("exists: ", exists)
		} else {
			fmt.Println("That value doesn't exist.")
			fmt.Println("val: ", val)
			fmt.Println("exists: ", exists)
		}

		fmt.Println(myGreeting)*/

	/*myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)
	delete(myGreeting, 7)
	fmt.Println(myGreeting)*/
	/*myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	// delete(myGreeting, 2)

	if val, exists := myGreeting[2]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)*/

	/*
		myGreeting := map[string]string{
			"Tim":   "Good morning!",
			"Jenny": "Bonjour!",
		}

		myGreeting["Harleen"] = "Howdy"

		fmt.Println(myGreeting)
		fmt.Println(len(myGreeting))
		myGreeting["Harleen"] = "Gidday"
		fmt.Println(myGreeting)
		delete(myGreeting, "Harleen")
		fmt.Println(myGreeting)*/

	/*	myGreeting := map[string]string{
			"Tim":   "Good morning!",
			"Jenny": "Bonjour!",
		}

		fmt.Println(myGreeting["Jenny"])*/
	/*myGreeting := map[string]string{}
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)*/
	/*myGreeting := make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)*/
	/*	var myGreeting = make(map[string]string)
		myGreeting["Tim"] = "Good morning."
		myGreeting["Jenny"] = "Bonjour."

		fmt.Println(myGreeting)*/
	/*var myGreeting map[string]string
	fmt.Println(myGreeting)
	fmt.Println(myGreeting == nil)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."*/
}

// add these lines:
/*
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
*/
// and you will get this:
// panic: assignment to entry in nil map
