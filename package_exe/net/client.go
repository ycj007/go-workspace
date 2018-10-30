package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strconv"
)

const (
	URL = "https://www.baidu.com/"
)

func GetContent(url string) string {

	connection, error := net.Dial("tcp", "localhost:8080")

	if error != nil {
		log.Fatal(error)
	}
	str, error := strconv.Atoi("\n")
	status, error := bufio.NewReader(connection).ReadString(byte(str))
	fmt.Println(status)
	data := []byte{}

	for index, _ := connection.Read(data); index < 0; {

		fmt.Println(string(data))
	}

	return status

}

func main() {
	GetContent(URL)

}
