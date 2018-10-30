package main

import (
	"fmt"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	fmt.Println(conn.RemoteAddr().String())

	conn.Write([]byte("heloword"))

}
