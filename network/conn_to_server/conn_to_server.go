package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	buffer := make([]byte, 1024)
	reader := bufio.NewReader(conn)

	n, err := reader.Read(buffer)
	if err != nil {
		fmt.Errorf("read failed: %v", err)
	}

	message := strings.ToUpper(string(buffer[:n]))
	fmt.Printf("Message: %v", message)
}
