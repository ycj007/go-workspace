package main

import "fmt"

func main() {
	test18()
}

func test18() {
	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 0, 4)
		for j := 0; j < 4; j++ {
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println(transactions)

}

func test17() {
	var records [][]string
	// student 1
	student1 := make([]string, 4)
	student1[0] = "Foster"
	student1[1] = "Nathan"
	student1[2] = "100.00"
	student1[3] = "74.00"
	// store the record
	records = append(records, student1)
	// student 2
	student2 := make([]string, 4)
	student2[0] = "Gomez"
	student2[1] = "Lisa"
	student2[2] = "92.00"
	student2[3] = "96.00"
	// store the record
	records = append(records, student2)
	// print
	fmt.Println(records)
}

func test16() {
	student := make([]string, 35)
	students := make([][]string, 35)
	student[0] = "Todd"
	// student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}

func test15() {
	var student []string
	var students [][]string
	//student[0] = "Todd"
	student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}
func test14() {
	student := []string{}
	students := [][]string{}
	//student[0] = "Todd"
	student = append(student, "Todd")
	fmt.Println(student)
	fmt.Println(students)
}

func test13() {
	student := make([]string, 35)
	students := make([][]string, 35)
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}

func test12() {
	var student []string
	var students [][]string
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

}
func test11() {
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)

}

func test10() {
	mySlice := []string{"hello", "1", "3", "5"}
	otherSlice := []string{"world", "23", "34234", "3242", "324"}

	mySlice = append(mySlice, otherSlice...)
	fmt.Println(mySlice)
	mySlice = append(mySlice[:2], mySlice[3:]...)
	fmt.Print(mySlice)
}

func test09() {
	mySlice := []string{"hello"}
	otherSlice := []string{"world"}

	mySlice = append(mySlice, otherSlice...)
	fmt.Print(mySlice)
}
func test08() {
	mySlice := []int{1, 2, 3, 4, 5}
	myOtherSlice := []int{6, 7, 8, 9}

	mySlice = append(mySlice, myOtherSlice...)

	fmt.Println(mySlice)
}

func test07() {
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Suprabadham")
	greeting = append(greeting, "Zǎo'ān")
	greeting = append(greeting, "Ohayou gozaimasu")
	greeting = append(greeting, "gidday")

	fmt.Println(greeting[6])
	fmt.Println(len(greeting))
	fmt.Println(cap(greeting))
}
func test06() {
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting = append(greeting, "Suprabadham")

	fmt.Println(greeting[3])
}

func test04() {

	customerNumber := make([]int, 3)
	// 3 is length & capacity
	// // length - number of elements referred to by the slice
	// // capacity - number of elements in the underlying array
	customerNumber[0] = 7
	customerNumber[1] = 10
	customerNumber[2] = 15

	fmt.Println(customerNumber[0])
	fmt.Println(customerNumber[1])
	fmt.Println(customerNumber[2])

	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array
	// you could also do it like this

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "dias!"

	fmt.Println(greeting[2])
}

func test05() {
	greeting := make([]string, 3, 5)
	// 3 is length - number of elements referred to by the slice
	// 5 is capacity - number of elements in the underlying array

	greeting[0] = "Good morning!"
	greeting[1] = "Bonjour!"
	greeting[2] = "buenos dias!"
	greeting[3] = "suprabadham"

	fmt.Println(greeting[2])
}

func test03() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	fmt.Print("[1:2] ")
	fmt.Println(greeting[1:2])
	fmt.Print("[:2] ")
	fmt.Println(greeting[:2])
	fmt.Print("[5:] ")
	fmt.Println(greeting[5:])
	fmt.Print("[:] ")
	fmt.Println(greeting[:])
}
func test02() {
	var results []int
	fmt.Println(results)

	mySlice := []string{"a", "b", "c", "g", "m", "z"}
	fmt.Println(mySlice)
	fmt.Println(mySlice[2:4])  // slicing a slice
	fmt.Println(mySlice[2])    // index access; accessing by index
	fmt.Println("myString"[2]) // index access; accessing by index
}
func test01() {
	greeting := []string{
		"Good morning!",
		"Bonjour!",
		"dias!",
		"Bongiorno!",
		"Ohayo!",
		"Selamat pagi!",
		"Gutten morgen!",
	}

	for i, currentEntry := range greeting {
		fmt.Println(i, currentEntry)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}
}
