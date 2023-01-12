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
	//init da connection
	conn, err := net.Dial("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Fatalf("Error encountered: %v", err)
	}
	wg.Add(1)
	go readmsg(conn)
	go writemsg(conn)
	wg.Wait()

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
		c.Write([]byte(fmt.Sprintf("\nClient says: %v", text)))
		if text == "quit" {
			wg.Done()
			return
		}
	}

}
