package get

import (
	"encoding/json"
	"fmt"
	"log"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Locations    []string   `json:"locations"`
	ConcertDates []string   `json:"concertDates"`
	Relations    string   `json:"relations"`
}



func GetArtists() []Artist {
	data := GetData("https://groupietrackers.herokuapp.com/api/artists")
	artists := []Artist{}
	
	err := json.Unmarshal(data, &artists)
	if err != nil {
		log.Panic("Problem in GetArtists function when unmarshalling data: ", err)
		return []Artist{}
	}
	
	fmt.Println(artists[0].Id)
	fmt.Println(artists[0].Image)
	fmt.Println(artists[0].Name)
	fmt.Println(artists[0].Members)
	fmt.Println(artists[0].CreationDate)
	fmt.Println(artists[0].FirstAlbum)
	fmt.Println(artists[0].Locations)
	fmt.Println(artists[0].ConcertDates)
	fmt.Println(artists[0].Relations)

	return artists
}
