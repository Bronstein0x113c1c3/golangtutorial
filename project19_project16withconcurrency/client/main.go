package main

import (
	"flag"
	"project16/implementation"
)

var (
	ip   = flag.String("ip", "0.0.0.0", "provide your ip here!")
	port = flag.Int("port", 1, "provide your port here!")
	dir  = flag.String("dir", "./workdir", "provide your working directory!")
)

func main() {
	flag.Parse()
	client_info := implementation.New(*ip, int16(*port), *dir)
	client_info.InitDaConnection()
	client_info.Transfer("test.pdf")
}
