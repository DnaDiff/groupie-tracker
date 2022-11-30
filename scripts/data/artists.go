package data

import (
	"encoding/json"
	"log"
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


func GetArtists() []Artist {
	data := GetData("https://groupietrackers.herokuapp.com/api/artists")
	artists := []Artist{}
	// handle error if data is nil
	err := json.Unmarshal(data, &artists)
	if err != nil {
		log.Panic("Problem in GetArtists function when unmarshalling data: ", err)
		return []Artist{}
	}
	return artists
}
