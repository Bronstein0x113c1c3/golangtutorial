package main

import (
	"client_transfer/implementation"
	_ "client_transfer/implementation"
	"fmt"
)

func main() {
	service := implementation.New("127.0.0.1", "9090")
	service.Init()
	service.ReturnSelf().Input = service.GetTheFile("test.jpg")
	status := service.Transfer()
	service.Exit()
	if status == true {
		fmt.Println("successful")
	}
}
