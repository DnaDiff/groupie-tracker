package getCocc

import (
	"encoding/json"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    string   `json:"locations"`
	ConcertDates string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}

// getData.GetArtists loads in all data from the groupietracker apis'
// first page
func Artists() []Artist {
	data := GetData("https://groupietrackers.herokuapp.com/api/artists")
	artists := []Artist{}
	json.Unmarshal(data, &artists)
	return artists
}
