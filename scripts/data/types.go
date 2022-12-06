package data

import (
	"encoding/json"
	"groupie-tracker/scripts/tools"
	"strings"
)

// Artist type
type rawArtist struct {
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

func getArtists() []rawArtist {
	artists := []rawArtist{}
	getData("https://groupietrackers.herokuapp.com/api/artists", &artists)
	return artists
}

// Dates type
type rawDates struct {
	Index []struct {
		ID    int      `json:"id"`
		Dates []string `json:"dates"`
	} `json:"index"`
}

func getDates() rawDates {
	dates := rawDates{}
	getData("https://groupietrackers.herokuapp.com/api/dates", &dates)
	return dates
}

// Locations type
type rawLocations struct {
	Index []struct {
		ID        int      `json:"id"`
		Locations []string `json:"locations"`
		Dates     string   `json:"dates"`
	} `json:"index"`
}

func getLocations() rawLocations {
	locations := rawLocations{}
	getData("https://groupietrackers.herokuapp.com/api/locations", &locations)
	return locations
}

// Relation type
type rawRelations struct {
	Index []struct {
		DatesLocations map[string][]string `json:"datesLocations"`
	} `json:"index"`
}

func getRelations() rawRelations {
	// Initalize to store the raw data
	relations := rawRelations{}
	getData("https://groupietrackers.herokuapp.com/api/relation", &relations)
	// Convert the names to the correct format
	// EX: "los_angeles-usa" -> "Los Angeles, USA"
	marshalledRelations, _ := json.Marshal(relations.Index)
	marshalledRelations = []byte(
		tools.Title(
			strings.ReplaceAll(
				strings.ReplaceAll(
					string(marshalledRelations), "_", " "), "-", ", ")))
	// Reset to avoid duplicates when unmarshalling
	relations = rawRelations{}
	json.Unmarshal(marshalledRelations, &relations.Index)

	// Convert the dates to the correct format
	for i := 0; i < len(relations.Index); i++ {
		for location, dates := range relations.Index[i].DatesLocations {
			// Convert the dates to the correct format
			// EX: "2019-01-01" -> "January 1, 2019"
			for dateIndex, date := range dates {
				relations.Index[i].DatesLocations[location][dateIndex] = tools.Date(date)
			}
		}
	}

	return relations
}
