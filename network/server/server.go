package main

import (
	"fmt"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		fmt.Println("Error listening:", err.Error())
	}
	defer listener.Close()
	fmt.Println("Listening on :8081")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		fmt.Println("Accepting connection from ", conn.RemoteAddr())
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()

	message := []string{"string1", "string2", "string3"}
	for _, msg := range message {
		fmt.Fprintln(c, msg)
	}

	fmt.Println("Messages sent")
}
