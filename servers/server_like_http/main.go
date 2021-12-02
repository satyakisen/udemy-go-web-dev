package main

import (
	"bufio"
	"fmt"
	"html/template"
	"log"
	"net"
	"strings"
)

func main() {

	listner, err := net.Listen("tcp", ":8082")

	if err != nil {
		log.Fatalln(err)
	}

	defer listner.Close()

	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	request(conn)
	response(conn)
}

func request(conn net.Conn) {
	sc := bufio.NewScanner(conn)
	i := 0
	for sc.Scan() {
		text := sc.Text()
		fmt.Println(text)
		if i == 0 {
			m := strings.Fields(text)[0]
			fmt.Println("***METHOD: ", m)
		}
		if text == "" {
			break
		}
		i++
	}
}

func response(conn net.Conn) {
	tpl, err := template.ParseFiles("tpl.gohtml")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	err = tpl.Execute(conn, "Ankita")

	if err != nil {
		log.Fatalln(err)
	}
}
