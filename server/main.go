package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	server, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := server.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	msg, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(msg)
}
