package main

import (
	"fmt"
	"net/http"
	api "project17/api_implementation"
	"project17/entity"
	"project17/implementation"
)

func main() {
	// set up an entity to operate first, just to make implementation valid (kind of a trick)
	e := &implementation.Basics{
		Entity: new(entity.StudentList),
	}
	//setup a server and then gather everything
	operate_serv := &api.Serv{
		Inner: e,
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/addstudent", operate_serv.AddHandler)
	mux.HandleFunc("/getstudent", operate_serv.GetHandler)
	mux.HandleFunc("/updatestudent", operate_serv.UpdateHandler)
	mux.HandleFunc("/deletestudent", operate_serv.DeleteHandler)
	fmt.Println("Server is listening in port 8080....")
	fmt.Println(http.ListenAndServe(":8080", mux))
}
