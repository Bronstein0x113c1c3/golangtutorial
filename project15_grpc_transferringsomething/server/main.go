package main

import (
	"fmt"
	"server_transfer/implementation"
)

func main() {
	server := implementation.New("192.168.243.131", "9090")
	fmt.Println("server is created")
	server.InitDaConnection()
}
