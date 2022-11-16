package main

import (
	"fmt"
	get "groupie-tracker/server/Data"
)

func main() {

	// server.StartServer()
	// print the first artist in the list with newlines between each field
	artists := get.GetArtists()
	fmt.Println(artists[0].Id)
	fmt.Println(artists[0].Image)
	fmt.Println(artists[0].Name)
	fmt.Println(artists[0].Members)
	fmt.Println(artists[0].CreationDate)
	fmt.Println(artists[0].FirstAlbum)
	fmt.Println(artists[0].Locations)
	fmt.Println(artists[0].ConcertDates)
	fmt.Println(artists[0].Relations)

}
