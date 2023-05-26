package main

import (
	"excercise4/mainwork"
	"flag"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var name = flag.String("name", "", "Enter the film you want to find. Use underscore \"_\" instead of space")

func main() {
	flag.Parse()
	if *name == "" {
		return
	}
	err := godotenv.Load()
	if err != nil {
		return
	}
	api_key := os.Getenv("API_KEY")
	fmt.Println(*name)
	mainwork.Find(*name, api_key)
}
