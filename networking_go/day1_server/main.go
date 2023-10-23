package main

import (
	"log"
	"net"
)

func handleconn(c net.Conn) {
	for {
		c.Write([]byte("Bonjour! \n"))
	}
	c.Close()
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, _ := lis.Accept()
		go handleconn(conn)
	}
}
