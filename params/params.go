package params

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func Echo() {

	fmt.Println(os.Args[0])
}

func EchoParamsEveryLine() {

	for index, value := range os.Args[1:] {

		fmt.Println(index)
		fmt.Println(value)
	}
}

func Uniq() {

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for index, i := range counts {
		fmt.Println(index)
		fmt.Println(i)
	}
}

func Uniq2() {

	counts := make(map[string]int)
	filename := os.Args[1]

	data, error := ioutil.ReadFile(filename)
	if error != nil {

	}

	fileContent := string(data)
	for _, line := range strings.Split(fileContent, "\n") {
		counts[line]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}
