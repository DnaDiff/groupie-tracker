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
	// relate := get.GetRelations()
	// fmt.Println(relate)
	// print the first artist in the list with newlines between each field
	// nice := 13
	// fmt.Println(relate.Index[1].Id)
	// fmt.Println(relate.Index[1].DatesLocations)

}
