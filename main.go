package main

import (
	"fmt"
	get "groupie-tracker/server/Data"
)

func main() {

	// server.StartServer()
	// print the first artist in the list with newlines between each field
	nice := 13
	artists := get.GetArtists()
	dates := get.GetDates()
	locations := get.GetLocations()

	// // check if the length of the dates and locations slices are the same as the length of the artists slice
	// fmt.Println(len(artists))
	// fmt.Println(len(dates.Index))
	// fmt.Println(len(locations.Index))

	fmt.Println(artists[nice].Id)
	fmt.Println(artists[nice].Image)
	fmt.Println(artists[nice].Name)
	fmt.Println(artists[nice].Members)
	fmt.Println(artists[nice].CreationDate)
	fmt.Println(artists[nice].FirstAlbum)
	fmt.Println(locations.Index[nice].Locations)
	fmt.Println(dates.Index[nice].Dates)
	fmt.Println(artists[nice].Relations)

	// everyone := get.Epicmerge()
	// fmt.Println(everyone[2])
}
