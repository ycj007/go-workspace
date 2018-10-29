package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, _ := http.Get("http://www.baidu.com")
	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Print(string(data))
	resp.Body.Close()
}
