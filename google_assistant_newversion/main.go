package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	authentication "project_a/authentication"
	serv "project_a/service_implementation"
	"runtime"
	"time"
)

func main() {
	//Step 1: Initialize the infomation and authentication
	a := serv.Assistant{
		Client_ID:     authentication.Client_ID,
		Client_Secret: authentication.Client_secret,
	}
	a.Init()
	fmt.Println(a.Authentication_Info.Token.)
	for {
		conversation, err := a.NewConversation(300 * time.Second)
		defer conversation.Assist_Client.CloseSend()
		if err != nil {
			panic(err)
		}

		fmt.Print("Query: ")
		query := readLine()
		if query != "quit" {
			resp := conversation.Query(query)
			fmt.Println("Response: ", resp.GetDialogStateOut().GetSupplementalDisplayText())
		} else {
			runtime.GC()
		}
	}

}
func readLine() string {
	rdr := bufio.NewReader(os.Stdin)
	line, err := rdr.ReadString('\n')
	switch err {
	case io.EOF:
		os.Exit(0)
	default:
		if err != nil {
			panic(err)
		}
	}
	return line
}
func pressEnter() {
	bufio.NewReader(os.Stdin).ReadLine()
}
