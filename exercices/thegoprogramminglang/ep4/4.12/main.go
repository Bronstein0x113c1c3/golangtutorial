package main

import (
	"ep4/work"
	"flag"
	"fmt"
	"log"
)

var ep *int = flag.Int("ep", 0, "provide the ep for the query")

func main() {
	flag.Parse()
	info, err := work.GetBook(*ep)
	if err != nil {
		log.Fatalln("Error encountered")
	}
	fmt.Println(info.Img)
}
