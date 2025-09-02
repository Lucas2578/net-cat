package main

import (
	"net-cat/server"
	"os"
)

func main() {
	args := os.Args
	var port string
	if len(args) == 1 {
		port = ":8989"
	} else {
		port = ":" + args[1]
	}

	server.Start(port)
}
