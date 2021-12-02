package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	lis, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Panic(err)
	}

	defer lis.Close()

	for {
		conn, err := lis.Accept()

		if err != nil {
			log.Panic(err)
		}

		go handleConnection(conn)

	}
}

func handleConnection(conn net.Conn) {
	err := conn.SetDeadline(time.Now().Add(10 * time.Second))

	if err != nil {
		fmt.Println("Conn Timed out")
	}

	scanner := bufio.NewScanner(conn)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		fmt.Fprintf(conn, "I heard you say: %s\n", line)
	}

	defer conn.Close()

	fmt.Println("*** CODE GOT HERE ***")

}
