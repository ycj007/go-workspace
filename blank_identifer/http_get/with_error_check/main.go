package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	resp, error := http.Get("http://www.baidu123.com")

	if error != nil {
		log.Fatal(error)

	}
	content, error := ioutil.ReadAll(resp.Body)
	if error != nil {
		log.Fatal(error)
	}
	fmt.Print(string(content))

}
