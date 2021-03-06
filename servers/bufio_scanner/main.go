package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer li.Close()
	for {
		conn, err := li.Accept()
		if err != nil {
			log.Panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	defer conn.Close()

	fmt.Println("Code got here.")
}
