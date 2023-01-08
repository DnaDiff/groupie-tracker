package main

import (
	"fmt"
	"groupie-tracker/scripts/server"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		server.Port = os.Args[1]
	} else if len(os.Args) > 2 {
		fmt.Println("Only input one argument to change port.")
		return
	}
	server.StartServer()
}
