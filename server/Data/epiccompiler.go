package get


import (
	"fmt"
)

func Epicmerge() {

	// GetArtists()
	// GetDates()
	// GetLocations()

	artists := GetArtists()
	dates := GetDates()
	locations := GetLocations()

	fmt.Println(artists)
	fmt.Println(dates)
	fmt.Println(locations)

	for i := 0; i < len(artists); i++ {
		artists[i].ConcertDates = dates.Index[i].Dates
		artists[i].Locations = locations.Index[i].Locations
	}

	fmt.Println(artists)

}
