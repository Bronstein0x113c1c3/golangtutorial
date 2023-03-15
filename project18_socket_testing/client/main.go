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
	//add waitgroup to hold the goroutines from being halted
	wg.Add(1)
	//split the connection into 2 distinct goroutines to make it run more efficiently
	go readmsg(conn)
	go writemsg(conn)
	//wait for the waitgroup to be done!
	wg.Wait()

}
func readmsg(c net.Conn) {
	//get it in while loop to hold it!
	for {
		b := make([]byte, 1024)
		c.Read(b)
		if len(b) != 0 {
			fmt.Println(string(b))
		}
	}
}
func writemsg(c net.Conn) {
	//also, but when we want to quit, just say "quit", the program halted!
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
