package main

import (
	"fmt"
	get "groupie-tracker/server/Data"
)

func main() {

	// server.StartServer()
	// print the first artist in the list with newlines between each field
	nice := 2
	artists := get.GetArtists()
	fmt.Println(artists[nice].Id)
	fmt.Println(artists[nice].Image)
	fmt.Println(artists[nice].Name)
	fmt.Println(artists[nice].Members)
	fmt.Println(artists[nice].CreationDate)
	fmt.Println(artists[nice].FirstAlbum)
	fmt.Println(artists[nice].Locations)
	fmt.Println(artists[nice].ConcertDates)
	fmt.Println(artists[nice].Relations)

}
