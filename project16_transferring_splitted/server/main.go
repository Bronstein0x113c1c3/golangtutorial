package main

import (
	"flag"
	service "project16/implementation"
)

func main() {
	ip := flag.String("ip", "0.0.0.0", "provide your ip here!")
	port := flag.Int("port", 1, "provide your port here!")
	dir := flag.String("dir", "./workdir", "provide your working directory!")
	server_info := service.New(*ip, int16(*port), *dir)
	server_info.InitDaConnection()
}
