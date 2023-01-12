package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	/*
		Step 1: Init the server, listen!
		Step 2: Accept client
		Step 3: Start chatting

		Idea, scenario:
			1st: one to one chatting server-client (single connections)
				1.1: One to one connection
				1.2: Many to one connection
				1.3: Isolating with each connection
			2nd: one client through one server to other clients:
			3rd: switching - one to one client!:
				- How to indicate the client by IP, MAC or generated ID....?
	*/
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("Error encountered: %v", err)
	}
	wg.Add(1)
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("Error encountered: %v", err)
		}
		go handler(conn)
		wg.Wait()
	}
}

// func readmsg(c net.Conn) {
// 	var b []byte
// 	c.Read(b)
// 	fmt.Println(string(b))
// }

//	func writemsg(c net.Conn) {
//		reader := bufio.NewReader(os.Stdin)
//		fmt.Print("Something to say: ")
//		text, _ := reader.ReadString('\n')
//		c.Write([]byte(fmt.Sprintf("Server says: %v", text)))
//	}
func handler(c net.Conn) {
	go readmsg(c)
	go writemsg(c)
}
func readmsg(c net.Conn) {
	for {
		b := make([]byte, 1024)

		c.Read(b)
		if len(b) != 0 {
			fmt.Println(string(b))
		}
	}
}
func writemsg(c net.Conn) {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		c.Write([]byte(fmt.Sprintf("\nServer says: %v", text)))
		if text == "quit" {
			wg.Done()
			return
		}
	}
}
